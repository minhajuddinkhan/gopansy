package router

import (
	"fmt"
	"net/http"

	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//AuthorizationBody AuthorizationBody
type AuthorizationBody struct {
	Authorization string
}

//SayHello say hello dummy func.
func SayHello(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(constants.Authorization)
	fmt.Println("USER!!", user)

	fmt.Fprintf(w, "VIP!")
}
