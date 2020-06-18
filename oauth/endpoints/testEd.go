package endpoints

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v4/server"
	"net/http"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 16:36
 * @Title:
 * --- --- ---
 * @Desc:
 */

type TestEndpoint struct {
	oauthServer *server.Server
}

func NewTestEndpoint(oauthServer *server.Server) *TestEndpoint {
	return &TestEndpoint{oauthServer}
}

func (e *TestEndpoint) Test(ctx *gin.Context) {
	token, err := e.oauthServer.ValidationBearerToken(ctx.Request)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusBadRequest)
		return
	}

	data := map[string]interface{}{
		"expires_in": int64(token.GetAccessCreateAt().Add(token.GetAccessExpiresIn()).Sub(time.Now()).Seconds()),
		"client_id":  token.GetClientID(),
		"user_id":    token.GetUserID(),
	}
	//e := json.NewEncoder(c.Writer)
	//e.SetIndent("", "  ")
	//e.Encode(data)
	ctx.JSON(200, &gin.H{
		"data": data,
	})
}
