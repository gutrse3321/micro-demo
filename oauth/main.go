package main

import (
	"github.com/dgrijalva/jwt-go"
	"gopkg.in/oauth2.v4/errors"
	"gopkg.in/oauth2.v4/generates"
	"gopkg.in/oauth2.v4/manage"
	"gopkg.in/oauth2.v4/models"
	"gopkg.in/oauth2.v4/server"
	"gopkg.in/oauth2.v4/store"
	"log"
	"net/http"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/1 16:40
 * @Title:
 * --- --- ---
 * @Desc:
 */

func main() {
	manager := manage.NewDefaultManager()
	manager.SetAuthorizeCodeTokenCfg(manage.DefaultPasswordTokenCfg)

	//redis token store
	//manager.MapTokenStorage(redisStore.NewRedisStore(&redis.Options{Addr: "207.148.105.249:6379", DB: 2}))
	tokenStore, _ := store.NewMemoryTokenStore()
	manager.MapTokenStorage(tokenStore)

	clientStore := store.NewClientStore()
	clientStore.Set("asuka", &models.Client{
		ID:     "asuka",
		Secret: "akiTomonori",
		Domain: "localhost",
	})
	manager.MapClientStorage(clientStore)

	//使用jwt加密
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate([]byte("akiTomonori"), jwt.SigningMethodHS512))

	srv := server.NewServer(server.NewConfig(), manager)
	srv.SetAllowGetAccessRequest(true)
	srv.SetClientInfoHandler(server.ClientFormHandler)
	//srv.SetPasswordAuthorizationHandler(func(username, password string) (userID string, err error) {
	//	if username == "admin" && password == "111111" {
	//		userID = "fucker"
	//	}
	//	return
	//})

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
		err := srv.HandleAuthorizeRequest(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	})

	http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
		srv.HandleTokenRequest(w, r)
	})
	log.Fatal(http.ListenAndServe(":9090", nil))
}
