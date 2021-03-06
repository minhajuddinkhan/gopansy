package router

import (
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/go-playground/validator.v9"

	"github.com/darahayes/go-boom"

	jwt "github.com/dgrijalva/jwt-go"
	dbsql "github.com/jmoiron/sqlx"
	conf "github.com/minhajuddinkhan/gopansy/config"
	constants "github.com/minhajuddinkhan/gopansy/constants"
	helpers "github.com/minhajuddinkhan/gopansy/helpers"
	"github.com/minhajuddinkhan/gopansy/models"
	schema "github.com/minhajuddinkhan/gopansy/schema"
	"golang.org/x/crypto/bcrypt"
)

//Login Login
func Login(w http.ResponseWriter, r *http.Request) {

	var payload schema.Login

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&payload)
	if err != nil {
		boom.BadData(w, "Cannot decode Request Body")
		return
	}

	login := schema.Login{
		Username: payload.Username,
		Password: payload.Password,
	}
	v := validator.New()
	err = login.Validate(v)
	if err != nil {
		boom.BadRequest(w, err.Error())
		return
	}
	db := r.Context().Value(constants.DbKey).(*dbsql.DB)

	type UserRole struct {
		models.User
		models.Role
	}
	userWithRole := UserRole{}

	user := models.User{
		Username: &payload.Username,
	}
	row := user.GetByUsername(db)
	err = row.StructScan(&userWithRole)
	if err != nil {
		boom.Internal(w)
		return
	}

	if userWithRole.Username == nil {
		boom.Unathorized(w, "Invalid Username or Password")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(*userWithRole.HashedPassword), []byte(payload.Password))
	if err != nil {
		boom.Unathorized(w, "Invalid Username or Password")
		return
	}

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": userWithRole.Name,
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
	})
	token, err := signer.SignedString([]byte(conf.GetConfig().Jwt.Secret))
	if err != nil {
		boom.BadImplementation(w, "Could not sign JWT token")
	}

	helpers.Respond(w, struct {
		Authroziation string `json:"Authorization"`
	}{
		token,
	})

}
