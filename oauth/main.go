package oauth

import (
	"crypto/rsa"
	"demo/oauth/cert"
	"encoding/json"
	"fmt"
	"github.com/RangelReale/osin"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/wire"
	"github.com/gutrse3321/aki/pkg/app"
	akiHttp "github.com/gutrse3321/aki/pkg/transports/http"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"net/url"
)

func NewOptions(v *viper.Viper, logger *zap.Logger) (*app.Options, error) {
	var err error

	opt := &app.Options{}
	if err = v.UnmarshalKey("app", opt); err != nil {
		return nil, errors.Wrap(err, "unmarshal app config error")
	}

	logger.Info("load application config success")
	return opt, err
}

func NewApp(opt *app.Options, logger *zap.Logger, httpServer *akiHttp.Server) (*app.Application, error) {
	application, err := app.New(opt, logger, app.HttpServerOption(httpServer))
	if err != nil {
		return nil, errors.Wrap(err, "new application error")
	}

	return application, nil
}

var WireSet = wire.NewSet(NewApp, NewOptions)

func main_old() {
	config := &osin.ServerConfig{
		AuthorizationExpiration:   250,
		AccessExpiration:          3600,
		TokenType:                 "Bearer",
		AllowedAuthorizeTypes:     osin.AllowedAuthorizeType{osin.CODE, osin.TOKEN},
		AllowedAccessTypes:        osin.AllowedAccessType{osin.AUTHORIZATION_CODE, osin.PASSWORD},
		ErrorStatusCode:           200,
		AllowClientSecretInParams: false,
		AllowGetAccessRequest:     true,
		RetainTokenAfterRefresh:   false,
	}
	server := osin.NewServer(config, NewTestStorage())

	var err error
	var accessTokenGenJWT AccessTokenGenJWT

	if accessTokenGenJWT.PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(cert.PrivatekeyPEM); err != nil {
		panic(err)
	}

	if accessTokenGenJWT.PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(cert.PublickeyPEM); err != nil {
		panic(err)
	}

	server.AccessTokenGen = &accessTokenGenJWT

	//授权请求 grant_type: 授权码 authorization_code 或者 账号密码 password
	//aauthorization_code endpoint
	//http://localhost:14000/auth?
	//response_type=code
	//&client_id=1234
	//&redirect_uri=http://（%3A%2F%2F）localhost%3A14000/（%2F）appauth/（%2F）code=生成的
	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		//handleAuthorizeRequest 授权码获取token会根据客户端设置的跳转链接去获取授权码
		if ar := server.HandleAuthorizeRequest(resp, r); ar != nil {
			if !handleLogin(ar, w, r) {
				return
			}

			ar.Authorized = true
			server.FinishAuthorizeRequest(resp, r, ar)
		}

		if resp.IsError && resp.InternalError != nil {
			fmt.Println("/auth error:", resp.InternalError)
		}
		osin.OutputJSON(resp, w, r)
	})

	// Access token endpoint
	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()

		if ar := server.HandleAccessRequest(resp, r); ar != nil {
			ar.Authorized = true
			server.FinishAccessRequest(resp, r, ar)
		}
		if resp.IsError && resp.InternalError != nil {
			fmt.Printf("ERROR: %s\n", resp.InternalError)
		}
		osin.OutputJSON(resp, w, r)
	})

	//查询当前令牌信息，eg：accesstoken refreshtoken 过期时间 clientid
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		resp := server.NewResponse()
		defer resp.Close()
		if ir := server.HandleInfoRequest(resp, r); ir != nil {
			server.FinishInfoRequest(resp, r, ir)
		}
		osin.OutputJSON(resp, w, r)
	})

	//生成授权码
	http.HandleFunc("/app/auth/code", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		code := r.FormValue("code")
		if code == "" {
			w.Write([]byte("Nothing to do"))
			return
		}

		jr := make(map[string]interface{})

		//build access code url
		aurl := fmt.Sprintf("/token?grant_type=authorization_code&client_id=1234&state=xyz&redirect_uri=%s&code=%s",
			url.QueryEscape("http://localhost:14000/app/auth/code"), url.QueryEscape(code))

		if r.FormValue("doparse") == "1" {
			err := downloadAccessToken(fmt.Sprintf("http://localhost:14000%s", aurl),
				&osin.BasicAuth{"1234", "aabbccdd"}, jr)
			if err != nil {
				w.Write([]byte(err.Error()))
			}
		}

		//show json error
		if erd, ok := jr["error"]; ok {
			w.Write([]byte(fmt.Sprintf("ERROR: %s<br/>\n", erd)))
		}

		// show json access token
		if at, ok := jr["access_token"]; ok {
			w.Write([]byte(fmt.Sprintf("ACCESS TOKEN: %s<br/>\n", at)))
		}

		w.Write([]byte(fmt.Sprintf("FULL RESULT: %+v<br/>\n", jr)))
		// output links
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Goto Token URL</a><br/>", aurl)))
		cururl := *r.URL
		curq := cururl.Query()
		curq.Add("doparse", "1")
		cururl.RawQuery = curq.Encode()
		w.Write([]byte(fmt.Sprintf("<a href=\"%s\">Download Token</a><br/>", cururl.String())))
	})

	http.ListenAndServe(":14000", nil)
}

func handleLogin(ar *osin.AuthorizeRequest, w http.ResponseWriter, r *http.Request) bool {
	r.ParseForm()
	if r.Method == "POST" && r.FormValue("login") == "tomo" && r.FormValue("password") == "tomo" {
		return true
	}
	return false
}

func downloadAccessToken(url string, auth *osin.BasicAuth, output map[string]interface{}) error {
	preq, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	if auth != nil {
		preq.SetBasicAuth(auth.Username, auth.Password)
	}

	pclient := &http.Client{}
	presp, err := pclient.Do(preq)
	if err != nil {
		return err
	}

	if presp.StatusCode != 200 {
		return errors.New("Invalid status code")
	}

	jdec := json.NewDecoder(presp.Body)
	err = jdec.Decode(&output)

	return err
}

type AccessTokenGenJWT struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

/*
generate JWT access token
*/
func (a *AccessTokenGenJWT) GenerateAccessToken(data *osin.AccessData, generateRefresh bool) (accessToken string, refreshToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"cid": data.Client.GetId(),
		"exp": data.ExpireAt().Unix(),
	})

	accessToken, err = token.SignedString(a.PrivateKey)
	if err != nil {
		return
	}

	if !generateRefresh {
		return
	}

	//生成jwt刷新token
	token = jwt.NewWithClaims(jwt.SigningMethodRS256, jwt.MapClaims{
		"cid": data.Client.GetId(),
	})

	refreshToken, err = token.SignedString(a.PrivateKey)
	if err != nil {
		return
	}

	return
}
