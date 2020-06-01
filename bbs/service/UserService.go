package service

import (
	provider "github.com/gutrse3321/aki-remote"
	"github.com/gutrse3321/aki/persit/remote"
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
	userRemote *provider.UCenterRemote
}

func NewUserService(userRemote *provider.UCenterRemote) IUserService {
	return &UserServiceImpl{userRemote}
}

func (u *UserServiceImpl) GetUserInfo() (interface{}, error) {
	str := "ch"
	codeRemote, err := u.userRemote.GetUserBaseInfo(&str)
	if err != nil {
		return nil, err
	}

	return remote.ResolveRemote(*codeRemote)
}
