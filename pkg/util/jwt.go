package util

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Openid string `json:"openid"`
	Id     string `json:"id"`
	jwt.StandardClaims
}

var jwtSecret = []byte("f3978b57-4679-44b2-9e96-7f750cb2a0a5")

func GenerateToken(openid, id string) (string, error) {
	var jwtExpire int64 = 86400
	nowTime := time.Now()
	claims := Claims{
		openid,
		id,
		jwt.StandardClaims{
			ExpiresAt: nowTime.Unix() + jwtExpire,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
