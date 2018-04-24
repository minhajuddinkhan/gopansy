package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	conf "github.com/minhajuddinkhan/gopansy/config"

	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//AuthorizationBody AuthorizationBody
type AuthorizationBody struct {
	Authorization string
}

//SayHello say hello dummy func.
func SayHello(w http.ResponseWriter, r *http.Request) {
	auth := r.Context().Value(constants.Authorization).(string)
	decoded, _ := jwt.Parse(auth, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(conf.GetConfig().Jwt.Secret), nil
	})
	fmt.Println(decoded)
	body := AuthorizationBody{
		Authorization: auth,
	}
	w.Header().Set("content-type", "application/json")
	authJSON, _ := json.Marshal(body)
	w.Write(authJSON)
}
