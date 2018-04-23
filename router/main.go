package router

import (
	"github.com/gorilla/mux"
)

//Initiate Initiate
func Initiate() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", SayHello)
	return mux

}
