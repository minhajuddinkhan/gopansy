package router

import (
	"net/http"

	"github.com/minhajuddinkhan/gopansy/helpers"

	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//SayHello say hello dummy func.
func SayHello(w http.ResponseWriter, r *http.Request) {

	user := r.Context().Value(constants.Authorization)
	helpers.Respond(w, user)
}
