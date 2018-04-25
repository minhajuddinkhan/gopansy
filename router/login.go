package router

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	dbsql "github.com/jmoiron/sqlx"
	conf "github.com/minhajuddinkhan/gopansy/config"
	constants "github.com/minhajuddinkhan/gopansy/constants"
	"github.com/minhajuddinkhan/gopansy/models"
	"golang.org/x/crypto/bcrypt"
)

//Login Login
func Login(w http.ResponseWriter, r *http.Request) {

	var loginPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	Users := []models.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&loginPayload)
	if err != nil {
		fmt.Println("Cannot decode")
	}

	const rounds int = 10
	db := r.Context().Value(constants.DbKey).(*dbsql.DB)
	err = db.Select(&Users, "SELECT * FROM users  WHERE username = $1", loginPayload.Username)
	if err != nil {
		fmt.Println("cant scan", err)
	}
	defer r.Body.Close()

	user := Users[0]

	h := []byte(user.HashedPassword.String)
	p := []byte(loginPayload.Password)
	err = bcrypt.CompareHashAndPassword(h, p)
	if err != nil {
		fmt.Println("not equal bro", err)
	}

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": "admin",
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
	})
	token, err := signer.SignedString([]byte(conf.GetConfig().Jwt.Secret))
	if err != nil {
		fmt.Println("ERROR", err)
	}

	w.Header().Set("content-type", "application/json")
	encoder := json.NewEncoder(w)
	err = encoder.Encode(struct {
		Authroziation string `json:"Authorization"`
	}{
		token,
	})
	if err != nil {
		fmt.Println("Cannot encode")
	}

}
