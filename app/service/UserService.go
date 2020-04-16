package service

import (
	"demo/relatives/ucenterRemote"
	"demo/ucenter/provider/param"
	"github.com/google/wire"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 18:38
 * @Title:
 * --- --- ---
 * @Desc:
 */
type IUserService interface {
	GetUserInfo() interface{}
}

type UserServiceImpl struct {
}

func NewUserService() IUserService {
	return &UserServiceImpl{}
}

func (u *UserServiceImpl) GetUserInfo() interface{} {
	return ucenterRemote.GetUserBaseInfoRemote(param.GetUserBaseInfoArgs{})
}

var ProviderSet = wire.NewSet(NewUserService)
