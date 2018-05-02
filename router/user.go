package router

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/minhajuddinkhan/gopansy/schema"
	"gopkg.in/go-playground/validator.v9"

	"github.com/minhajuddinkhan/gopansy/helpers"

	"github.com/darahayes/go-boom"

	"github.com/gorilla/mux"
	sql "github.com/jmoiron/sqlx"
	"github.com/minhajuddinkhan/gopansy/constants"
	models "github.com/minhajuddinkhan/gopansy/models"
)

//GetUsers GetUsers
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value(constants.DbKey).(*sql.DB)

	type UserRole struct {
		models.User
		models.Role
	}
	userWithRoles := []UserRole{}

	u := models.User{}
	rows, err := u.GetAllUserWithRoles(db)
	if err != nil {
		boom.Internal(w)
		return
	}
	for rows.Next() {
		var r UserRole
		err := rows.StructScan(&r)
		if err != nil {
			boom.Internal(w)
			return
		}
		userWithRoles = append(userWithRoles, r)
	}

	helpers.Respond(w, userWithRoles)
}

//GetUserByID GetUserByID
func GetUserByID(w http.ResponseWriter, r *http.Request) {

	userID := mux.Vars(r)["id"]
	if len(userID) == 0 {
		boom.BadRequest(w, "User Id Required")
	}
	db := r.Context().Value(constants.DbKey).(*sql.DB)
	type UserRole struct {
		models.User
		models.Role
	}
	user := UserRole{}

	row := user.GetUserByID(db, userID)
	row.StructScan(&user)

	if len(*user.Username) == 0 {
		boom.NotFound(w, "User not found")
		return
	}

	helpers.Respond(w, user)
}

//CreateUser CreateUser
func CreateUser(w http.ResponseWriter, r *http.Request) {

	//_ := r.Context().Value(constants.DbKey).(*sql.DB)
	decoder := json.NewDecoder(r.Body)
	user := models.User{}
	decoder.Decode(&user)

	v := validator.New()
	err := v.Struct(schema.User{
		user.Username,
		user.Email,
		user.Password,
		user.PermitOneAllowed,
		user.PermitTwoAllowed,
		user.RoleID,
	})
	if err != nil {
		boom.BadRequest(w, err.Error())
		return
	}

	db := r.Context().Value(constants.DbKey).(*sql.DB)
	row := db.QueryRowx("SELECT u.* FROM users u WHERE u.username = $1 OR u.email = $2", user.Username, user.Email)
	row.StructScan(&user)

	if len(*user.ID) != 0 {
		boom.Conflict(w, "User with this username/email already exists")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(*user.Password), constants.HashRounds)

	var hashStr *string
	c := string(hash)
	hashStr = &c
	user.HashedPassword = hashStr

	_, err = user.CreateUser(db)

	if err != nil {
		boom.Internal(w)
		return
	}

	helpers.Respond(w, struct {
		Success bool `json:"success"`
	}{true})

}
