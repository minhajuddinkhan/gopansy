package router

import (
	dbsql "database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"

	constants "github.com/minhajuddinkhan/gopansy/constants"
)

//Login Login
func Login(w http.ResponseWriter, r *http.Request) {

	var loginPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginPayload)
	if err != nil {
		fmt.Println("Cannot decode")
	}
	db := r.Context().Value(constants.DbKey).(*dbsql.DB)
	result := db.QueryRow("SELECT * FROM users WHERE username = $1", loginPayload.Username)
	defer r.Body.Close()

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(loginPayload)
	if err != nil {
		fmt.Println("Cannot encode")
	}

}
