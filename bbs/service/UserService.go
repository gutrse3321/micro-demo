package service

import (
	"demo/relatives"
	"demo/ucenter/provider/param"
	"github.com/gutrse3321/aki-persit/remote"
)

/**
 * @Author: Tomonori
 * @Date: 2020/4/16 18:38
 * @Title:
 * --- --- ---
 * @Desc:
 */
type IUserService interface {
	GetUserInfo() (interface{}, error)
}

type UserServiceImpl struct {
	userRemote relatives.UCenterRemote
}

func NewUserService(userRemote relatives.UCenterRemote) IUserService {
	return &UserServiceImpl{userRemote}
}

func (u *UserServiceImpl) GetUserInfo() (interface{}, error) {
	codeRemote, err := u.userRemote.GetUserBaseInfo(param.GetUserBaseInfoArgs{})
	if err != nil {
		return nil, err
	}

	return remote.ResolveRemote(*codeRemote)
}
