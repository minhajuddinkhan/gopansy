package router

import (
	"github.com/gorilla/mux"
)

//Initiate Initiate
func Initiate() *mux.Router {
	mux := mux.NewRouter()
	mux.HandleFunc("/", SayHello)
	mux.HandleFunc("/login", Login).Methods("POST")

	mux.HandleFunc("/user/{id}", GetUserByID).Methods("GET")
	mux.HandleFunc("/user", GetUsers).Methods("GET")
	mux.HandleFunc("/user", CreateUser).Methods("POST")

	mux.HandleFunc("/forms/permitone", CreatePermitOneForm).Methods("POST")

	return mux

}
