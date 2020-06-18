package endpoints

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/oauth2.v4/server"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/10 16:26
 * @Title:
 * --- --- ---
 * @Desc:
 */

type LoginEndpoint struct {
	oauthServer *server.Server
}

func NewLoginEndpoint(oauthServer *server.Server) *LoginEndpoint {
	return &LoginEndpoint{oauthServer}
}

func (e *LoginEndpoint) Login(ctx *gin.Context) {
	err := e.oauthServer.HandleTokenRequest(ctx.Writer, ctx.Request)
	if err != nil {
		http.Error(ctx.Writer, err.Error(), http.StatusInternalServerError)
	}
}
