package service

import "github.com/gutrse3321/aki/persit/dto/user"

/**
 * @Author: Tomonori
 * @Date: 2020/4/15 17:51
 * @Title:
 * --- --- ---
 * @Desc:
 */
func GetUserBaseInfo() *user.UserBaseInfo {
	return &user.UserBaseInfo{
		Uid:         114514,
		Phone:       "18980165759",
		Email:       "hakurei@reimu.ru",
		RealName:    "tomonori",
		NickName:    "b32r",
		AvatarUri:   "",
		Description: "",
		Gender:      1,
		CreatedAt:   114514,
		UpdatedAt:   1919,
		DeletedAt:   0,
		Status:      1,
	}
}
