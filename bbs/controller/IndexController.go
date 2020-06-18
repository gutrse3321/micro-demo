package controller

import (
	"demo/bbs/service"
	"github.com/gin-gonic/gin"
	httpServer "github.com/gutrse3321/aki/pkg/transports/http"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 18:23
 * @Title:
 * --- --- ---
 * @Desc:
 */
type IndexController struct {
	userService service.IUserService
}

func NewIndexController(service service.IUserService) *IndexController {
	return &IndexController{userService: service}
}

func CreateInitControllersFn(index *IndexController) httpServer.InitControllers {
	return func(r *gin.Engine) {
		r.POST("/getUser", index.GetUserBaseInfo)
	}
}

func (i *IndexController) GetUserBaseInfo(ctx *gin.Context) {
	userInfo, err := i.userService.GetUserInfo()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	} else {
		ctx.JSON(http.StatusOK, &gin.H{
			"data": &userInfo,
		})
	}
}
