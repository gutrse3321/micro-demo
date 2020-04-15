package controller

import (
	"demo/relatives/ucenterRemote"
	"demo/ucenter/provider/param"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 18:23
 * @Title:
 * --- --- ---
 * @Desc:
 */
type Index struct {
}

func (i *Index) GetUserBaseInfo(ctx *gin.Context) {
	userInfo := ucenterRemote.GetUserBaseInfoRemote(param.GetUserBaseInfoArgs{})
	ctx.JSON(http.StatusOK, &userInfo)
}
