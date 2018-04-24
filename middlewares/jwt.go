package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//ParseJwt ParseJwt
func ParseJwt(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		// "role": "admin",
		// "exp":  time.Now().Add(time.Minute * 20).Unix(),
	})

	token, err := signer.SignedString("123")
	if err != nil {
		fmt.Println("ERROR", err)
	}
	ctx := context.WithValue(r.Context(), constants.Authorization, token)
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)

}
