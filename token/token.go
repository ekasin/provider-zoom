package token

import (
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"os"
)

var jwtKey = []byte("Qk2vHwnMdT0K1dXqWoFyYkMwt2CDkxOwYixV")

type Claims struct {
	jwt.StandardClaims
}



func TokenGenerate(apiSecret string, apiKey string) string {

	tknStr := os.Getenv("ZOOM_TOKEN")
	claims := &Claims{}
	tkn, _ := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {

		return jwtKey, nil
	})
	if tkn.Valid {
		return tknStr
	}
	apikey := "lNGJBHjuROOFKCM68LjH0g"
	expirationTime := time.Now().Add(20 * time.Minute)
	claims = &Claims{

		StandardClaims: jwt.StandardClaims{
		
			Audience:  " ",
			Issuer:    apikey,
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(jwtKey)
	fmt.Printf("\n New Token")
	return tokenString
	
}


func Refresh(previoustoken string) {
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
