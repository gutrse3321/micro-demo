package auth

import (
	"gopkg.in/oauth2.v4/errors"
	"gopkg.in/oauth2.v4/manage"
	"gopkg.in/oauth2.v4/server"
	"log"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 16:28
 * @Title:
 * --- --- ---
 * @Desc:
 */

func NewAuthorizationServer(manager *manage.Manager) *server.Server {
	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetAllowGetAccessRequest(false)
	srv.SetClientInfoHandler(server.ClientFormHandler)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("internal error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("response error:", re.Error.Error())
	})

	srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
		if username == "tomo" && password == "tomo" {
			userID = "114514"
		}
		return
	})

	return srv
}
