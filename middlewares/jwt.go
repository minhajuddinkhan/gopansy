package middlewares

import (
	"context"
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//ParseJwt ParseJwt
func ParseJwt(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	signer := jwt.New(jwt.SigningMethodHS256)

	token, err := signer.SignedString(constants.JwtSecret)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error while signing the token")
	}
	ctx := context.WithValue(r.Context(), constants.Authorization, token)
	r = r.WithContext(ctx)
	next.ServeHTTP(w, r)

}
