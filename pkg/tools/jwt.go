package tools

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"

	"go-gin-demo/pkg/setting"
)

var jwtSecret = []byte(setting.ENV.JwtSecret) //Secret，產生Token需要

type MyClaims struct { //自定義資料結構，需帶入官方指定物件
	Username             string `json:"username"`
	Password             string `json:"password"`
	jwt.RegisteredClaims        //官方指定
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour) //假定逾期3hour

	claims := MyClaims{
		username,
		password,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    "gin-blog",
		},
	} //自訂Claim內容

	//SigningMethodHS256、SigningMethodHS384、SigningMethodHS512三种crypto.Hash方案
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) //指定演算法，並將自訂Claim內容帶入求出Claim物件
	token, err := tokenClaims.SignedString(jwtSecret)                //利用Claim物件配合secret算出Token

	return token, err //回傳Token
}

func ParseToken(token string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil //依據官方範例這裡要回傳secret
	}) //將Token帶入解碼，注意中間&MyClaims{}資料結構要與NewWithClaims使用的Claims相同

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
			/*
				如果一切正常，此Claim會是完整的解碼結構
				(claims.Username 、claims.Password 、claims.RegisteredClaims.Issuer ...)

				type MyClaims struct {
					Username             string `json:"username"`
					Password             string `json:"password"`
					jwt.RegisteredClaims
				}
			*/
		}
	}

	return nil, err
}
