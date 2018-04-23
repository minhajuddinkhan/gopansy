package router

import (
	"fmt"
	"net/http"

	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//SayHello say hello dummy func.
func SayHello(w http.ResponseWriter, r *http.Request) {
	postgres := r.Context().Value(constants.DbKey)
	fmt.Println(postgres)
	fmt.Fprintf(w, "Hello World!")
}
