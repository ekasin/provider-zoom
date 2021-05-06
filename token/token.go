package token

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
}

func TokenGenerate(apisecret string, apikey string) string {
	var jwtKey = []byte(apisecret)
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &Claims{

		StandardClaims: jwt.StandardClaims{
			Audience:  " ",
			Issuer: apikey,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)

	fmt.Printf(tokenString)
	return tokenString
}

func Refresh(previoustoken string, apisecret string) {
	var jwtKey = []byte(apisecret)
	tknStr := previoustoken
	claims := &Claims{}
	tkn, _ := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil
	})
	if !tkn.Valid {
		fmt.Printf("token experied or unauthorised")
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().Unix()
	claims.Audience = " "
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		fmt.Printf("server Error")
		return
	}
	fmt.Printf(tokenString)

}
