package common

import (
	"gin_vue_bs/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//定义 jwt 加密密钥
var jwtKey = []byte("a_secret_crect")

//Claims .
type Claims struct {
	UserID uint
	jwt.StandardClaims
}

//ReleaseToken 登录成功则调用该方法发放token.
func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour) //token的过期时间7天
	claims := &Claims{
		UserID: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), //token发放时间
			Issuer:    "samtake",         //是谁发放的token
			Subject:   "user token",      //主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//ParseToken 解析token.
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
