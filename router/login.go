package router

import (
	dbsql "database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	constants "github.com/minhajuddinkhan/gopansy/constants"
	"github.com/minhajuddinkhan/gopansy/models"
)

//Login Login
func Login(w http.ResponseWriter, r *http.Request) {

	var loginPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var User models.User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginPayload)
	if err != nil {
		fmt.Println("Cannot decode")
	}
	fmt.Println("username", loginPayload.Username)
	db := r.Context().Value(constants.DbKey).(*dbsql.DB)
	result := db.QueryRow("SELECT username FROM users WHERE username = $1", loginPayload.Username)
	if err != nil {
		fmt.Println("Can't Scan Rows", err)
	}

	err = result.Scan(&User.Username)
	if err != nil {
		fmt.Println("cant scan", err)
	}
	fmt.Println("user!!!", User.Username)

	defer r.Body.Close()

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(loginPayload)
	if err != nil {
		fmt.Println("Cannot encode")
	}

}
