package router

import (
	_ "database/sql"
	"encoding/json"
	"net/http"
	"time"

	"gopkg.in/go-playground/validator.v9"

	"github.com/darahayes/go-boom"

	jwt "github.com/dgrijalva/jwt-go"
	dbsql "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	conf "github.com/minhajuddinkhan/gopansy/config"
	constants "github.com/minhajuddinkhan/gopansy/constants"
	helpers "github.com/minhajuddinkhan/gopansy/helpers"
	"github.com/minhajuddinkhan/gopansy/models"
	schema "github.com/minhajuddinkhan/gopansy/schema"
	"golang.org/x/crypto/bcrypt"
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
		boom.BadData(w, "Cannot decode Request Body")
		return
	}

	v := validator.New()
	err = v.Struct(schema.Login{
		loginPayload.Username,
		loginPayload.Password,
	})
	if err != nil {
		boom.BadRequest(w, err.Error())
		return
	}
	db := r.Context().Value(constants.DbKey).(*dbsql.DB)

	row := db.QueryRowx(`select u.*, r.name as rolename  
		from users u join roles r on (r.id = u.roleid)  
		where u.username = $1`, loginPayload.Username)

	var user models.User

	row.StructScan(&user)

	if user.ID == nil {
		boom.Unathorized(w, "Invalid Username or Password")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(*user.HashedPassword), []byte(loginPayload.Password))
	if err != nil {
		boom.Unathorized(w, "Invalid Username or Password")
		return
	}

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"role": user.RoleName,
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
