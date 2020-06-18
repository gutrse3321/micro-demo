package myJwt

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"gopkg.in/oauth2.v4"
)

/**
 * @Author: Tomonori
 * @Date: 2020/6/18 15:29
 * @Title:
 * --- --- ---
 * @Desc:
 */

type JWTGenerator struct {
	SignedKey []byte
}

type Claims struct {
	jwt.StandardClaims
	Uid      string `json:"uid"`
	ClientId string `json:"clientId"`
}

func (j *JWTGenerator) Token(ctx context.Context, data *oauth2.GenerateBasic, isGenRefresh bool) (access, refresh string, err error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Audience:  data.Client.GetID(),
			Subject:   data.UserID,
			ExpiresAt: data.TokenInfo.GetAccessCreateAt().Add(data.TokenInfo.GetAccessExpiresIn()).Unix(),
		},
		Uid:      data.UserID,
		ClientId: data.Client.GetID(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	access, err = token.SignedString(j.SignedKey)
	if err != nil {
		return "", "", err
	}

	if isGenRefresh {
		sign := uuid.Must(uuid.NewRandom()).String()
		refresh, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(sign))
		if err != nil {
			return "", "", err
		}
	}

	return
}
