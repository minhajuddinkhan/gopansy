package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/darahayes/go-boom"
	jwt "github.com/dgrijalva/jwt-go"
	conf "github.com/minhajuddinkhan/gopansy/config"
	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//EncodeJWT EncodeJWT
func EncodeJWT(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.RequestURI == "/login" {
		next.ServeHTTP(w, r)
		return
	}
	auth := r.Header.Get("Authorization")
	if len(auth) == 0 {
		boom.Unathorized(w, "Authorization Header not found")
		return
	}

	decoded, err := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(conf.GetConfig().Jwt.Secret), nil
	})
	if err != nil {
		boom.Unathorized(w, "Invalid Authorization Token")
		return
	}

	ctx := context.WithValue(r.Context(), constants.Authorization, decoded)
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)

}
