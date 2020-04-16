package controller

import (
	"demo/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 18:23
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Controller interface {
}

type Index struct {
	userService service.IUserService
}

func NewIndexController(service service.IUserService) *Index {
	return &Index{userService: service}
}

func (i *Index) GetUserBaseInfo(ctx *gin.Context) {
	userInfo := i.userService.GetUserInfo()
	ctx.JSON(http.StatusOK, &userInfo)
}

var ProviderSet = wire.NewSet(NewIndexController)
