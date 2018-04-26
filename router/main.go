package router

import (
	"github.com/gorilla/mux"
)

//Initiate Initiate
func Initiate() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", SayHello)
	mux.HandleFunc("/login", Login).Methods("POST")

	mux.HandleFunc("/user/{id}", GetUserById).Methods("GET")
	mux.HandleFunc("/user", GetUsers).Methods("GET")
	mux.HandleFunc("/user", CreateUser).Methods("POST")

	return mux

}
