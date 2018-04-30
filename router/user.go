package router

import (
	"encoding/json"
	"fmt"
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
	rows, err := db.Queryx("SELECT * FROM users u JOIN roles r on (u.roleId = r.id)")
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

func GetUserById(w http.ResponseWriter, r *http.Request) {

	userID := mux.Vars(r)["id"]
	if len(userID) == 0 {
		boom.BadRequest(w, "User Id Required")
	}
	db := r.Context().Value(constants.DbKey).(*sql.DB)
	type UserRole struct {
		models.User
		models.Role
	}

	row := db.QueryRowx("SELECT * from users u JOIN roles r on (u.roleId = r.id) WHERE u.id = $1", userID)
	user := UserRole{}
	row.StructScan(&user)

	if len(*user.Username) == 0 {
		boom.NotFound(w, "User not found")
		return
	}

	helpers.Respond(w, user)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {

	//_ := r.Context().Value(constants.DbKey).(*sql.DB)
	decoder := json.NewDecoder(r.Body)
	user := models.UserCreateRequest{}
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

	u := models.User{}

	db := r.Context().Value(constants.DbKey).(*sql.DB)
	row := db.QueryRowx("SELECT u.* FROM users u WHERE u.username = $1 OR u.email = $2", user.Username, user.Email)
	row.StructScan(&u)
	fmt.Println("*u.ID", *u.ID)
	if len(*u.ID) != 0 {
		boom.Conflict(w, "User with this username/email already exists")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(*user.Password), constants.HashRounds)
	db.Exec(`INSERT into users 
		(username, email, hashedPassword, permitOneAllowed, permitTwoAllowed, roleId) VALUES
		($1, $2, $3, $4, $5, $6) RETURNING id`, user.Username, user.Email, hash, user.PermitOneAllowed, user.PermitTwoAllowed, user.RoleID)

	if err != nil {
		boom.Internal(w)
		return
	}

	helpers.Respond(w, struct {
		Success bool `json:"success"`
	}{true})

}
