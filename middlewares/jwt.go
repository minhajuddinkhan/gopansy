package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	conf "github.com/minhajuddinkhan/gopansy/config"
	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//EncodeJWT EncodeJWT
func EncodeJWT(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": "admin",
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
	})
	token, err := signer.SignedString([]byte(conf.GetConfig().Jwt.Secret))
	if err != nil {
		fmt.Println("ERROR", err)
	}
	ctx := context.WithValue(r.Context(), constants.Authorization, token)
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)

}
