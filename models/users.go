package models

import (
	"database/sql"

	sqlx "github.com/jmoiron/sqlx"
	validator "gopkg.in/go-playground/validator.v9"
)

//User User
type User struct {
	ID               *string `json:"id"`
	Username         *string `json:"username" validate:"required"`
	Email            *string `json:"email" validate:"required"`
	HashedPassword   *string `json:"-"`
	Password         *string `json:"password,omitempty" validate:"required"`
	RoleID           *int    `json:"roleId, omitempty" validate:"required"`
	PermitOneAllowed *bool   `json:"permitOneAllowed" validate:"required"`
	PermitTwoAllowed *bool   `json:"permitTwoAllowed" validate:"required"`
}

//GetByUsername  GetByUsername
func (user *User) GetByUsername(db *sqlx.DB) *sqlx.Row {
	return db.QueryRowx(`select *  
		from users u join roles r on (r.id = u.roleId)  
		where u.username = $1`, user.Username)
}

//CreateUser CreateUser
func (user *User) CreateUser(db *sqlx.DB) (sql.Result, error) {

	return db.Exec(`INSERT into users 
		(username, email, hashedPassword, permitOneAllowed, permitTwoAllowed, roleId) VALUES
		($1, $2, $3, $4, $5, $6) RETURNING id`,
		user.Username,
		user.Email,
		user.HashedPassword,
		user.PermitOneAllowed,
		user.PermitTwoAllowed,
		user.RoleID)

}

//Validate Validate
func (user *User) Validate(v *validator.Validate) error {
	return v.Struct(user)
}

//GetByEmailOrUsername GetByEmailOrUsername
func (user *User) GetByEmailOrUsername(db *sqlx.DB) *sqlx.Row {
	return db.QueryRowx("SELECT u.* FROM users u WHERE u.username = $1 OR u.email = $2", user.Username, user.Email)

}

//GetUserByID GetUserByID
func (user *User) GetUserByID(db *sqlx.DB, userID string) *sqlx.Row {
	return db.QueryRowx("SELECT * from users u JOIN roles r on (u.roleId = r.id) WHERE u.id = $1", userID)
}

//GetAllUserWithRoles GetAllUserWithRoles
func (user *User) GetAllUserWithRoles(db *sqlx.DB) (*sqlx.Rows, error) {

	return db.Queryx("SELECT * FROM users u JOIN roles r on (u.roleId = r.id)")

}
