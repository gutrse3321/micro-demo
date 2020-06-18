package auth

import (
	"demo/oauth/myJwt"
	"github.com/go-redis/redis"
	"gopkg.in/oauth2.v4/manage"
	"gopkg.in/oauth2.v4/models"
	"gopkg.in/oauth2.v4/store"
	"time"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 15:40
 * @Title:
 * --- --- ---
 * @Desc:
 */

const JWTSign = "akiGate"

func NewOAuthManager(opt *OAuthOptions) *manage.Manager {
	manager := manage.NewDefaultManager()
	manager.SetPasswordTokenCfg(&manage.Config{
		AccessTokenExp:    opt.AccessTokenExp * time.Hour,
		RefreshTokenExp:   opt.RefreshTokenExp * time.Hour,
		IsGenerateRefresh: true,
	})

	//use redis token store
	manager.MapTokenStorage(NewRedisStore(&redis.Options{
		Addr:     opt.RedisOptions.Addr,
		Password: opt.RedisOptions.Password,
		DB:       opt.RedisOptions.Db,
	}))

	//use my jwt generator
	manager.MapAccessGenerate(&myJwt.JWTGenerator{SignedKey: []byte(JWTSign)})

	//set oauth clients
	clientStore := store.NewClientStore()
	for _, client := range opt.Clients {
		clientStore.Set(client.Id, &models.Client{
			ID:     client.Id,
			Secret: client.Secret,
		})
	}
	manager.MapClientStorage(clientStore)

	return manager
}
