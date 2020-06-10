package endpoints

import (
	"github.com/gin-gonic/gin"
	httpServer "github.com/gutrse3321/aki/pkg/transports/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/10 16:26
 * @Title:
 * --- --- ---
 * @Desc:
 */

type LoginEndpoint struct {
}

func NewLoginEndpoint() *LoginEndpoint {
	return &LoginEndpoint{}
}

func CreateInitControllersFn(login *LoginEndpoint) httpServer.InitControllers {
	return func(r *gin.Engine) {
		r.POST("/oauth/login", login.Login)
	}
}

func (e *LoginEndpoint) Login(ctx *gin.Context) {

}
