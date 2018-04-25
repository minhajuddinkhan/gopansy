package router

import (
	"github.com/gorilla/mux"
)

//Initiate Initiate
func Initiate() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", SayHello)
	mux.HandleFunc("/login", Login).Methods("POST")
	return mux

}
