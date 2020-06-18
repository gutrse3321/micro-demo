package auth

import "github.com/google/wire"

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 16:38
 * @Title:
 * --- --- ---
 * @Desc:
 */

var WireSet = wire.NewSet(
	NewOAuthOptions,
	NewOAuthManager,
	NewAuthorizationServer,
)
