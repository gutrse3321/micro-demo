package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	httpServer "github.com/gutrse3321/aki/pkg/transports/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/10 16:29
 * @Title:
 * --- --- ---
 * @Desc:
 */

func CreateInitControllersFn(login *LoginEndpoint, test *TestEndpoint) httpServer.InitControllers {
	return func(r *gin.Engine) {
		//LoginEndpoint
		r.POST("/oauth/login", login.Login)

		//TestEndpoint
		r.POST("/oauth/test", test.Test)
	}
}

var WireSet = wire.NewSet(
	NewLoginEndpoint,
	NewTestEndpoint,
	CreateInitControllersFn,
)
