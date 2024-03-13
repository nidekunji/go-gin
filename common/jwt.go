package common

import (
	"ginEssential/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId             uint
	jwt.StandardClaims //匿名字段
}

func ReleaseToken(user model.User) (string, error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "kunji",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 生成签名的Token字符串
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjIsImV4cCI6MTcxMDk2ODIyMywiaWF0IjoxNzEwMzYzNDIzLCJpc3MiOiJrdW5qaSIsInN1YiI6InVzZXIgdG9rZW4ifQ.EZ9qiGZH5hbFTSzbjuKnLXWfvN-lL2hgVRJkpRKACj0"
// 第一部分是header 第二部分是中间的信息 第三部分是1+2的hash值
// echo eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9 | base64 -D
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	// 创建Claims类型的实例
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, err
	})
	return token, claims, err
}
