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

	var loginPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	user := models.User{}

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
	const rounds int = 10
	db := r.Context().Value(constants.DbKey).(*dbsql.DB)

	row := db.QueryRowx(`select u.*, r.name as rolename  
		from users u join roles r on (r.id = u.roleid)  
		where u.username = $1`, loginPayload.Username)
	defer r.Body.Close()

	row.StructScan(&user)

	if len(user.ID.String) == 0 {
		boom.Unathorized(w, "Invalid Username or Password")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword.String), []byte(loginPayload.Password))
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
