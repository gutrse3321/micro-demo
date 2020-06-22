package auth

import (
	"gopkg.in/oauth2.v4"
	"gopkg.in/oauth2.v4/errors"
	"gopkg.in/oauth2.v4/manage"
	"gopkg.in/oauth2.v4/server"
	"log"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 16:28
 * @Title:
 * --- --- ---
 * @Desc:
 */
//&server.Config{
//TokenType:            "Bearer",
//AllowedResponseTypes: []oauth2.ResponseType{oauth2.Code, oauth2.Token},
//AllowedGrantTypes: []oauth2.GrantType{
//oauth2.AuthorizationCode,
//oauth2.PasswordCredentials,
//oauth2.ClientCredentials,
//oauth2.Refreshing,
//}

func NewAuthorizationServer(manager *manage.Manager) *server.Server {
	//利用server的SetAllowedGrantType方法来验证自定义的grantType
	//再想如何实现处理自定义的grantType
	srv := server.NewServer(&server.Config{
		TokenType:            "Bearer",
		AllowedResponseTypes: []oauth2.ResponseType{oauth2.Token},
		AllowedGrantTypes: []oauth2.GrantType{
			oauth2.PasswordCredentials,
			oauth2.Refreshing,
		},
	}, manager)
	srv.SetAllowGetAccessRequest(false)
	srv.SetClientInfoHandler(server.ClientBasicHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("internal error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("response error:", re.Error.Error())
	})

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "tomo" && password == "tomo" {
			userID = "114514"
		}
		return
	})

	return srv
}
