package models

import (
	"database/sql"
)

//User User
type User struct {
	ID               sql.NullString `json:"id"`
	Username         sql.NullString `json:"username"`
	Email            sql.NullString `json:"email"`
	HashedPassword   sql.NullString `json:"hashedPassword"`
	RoleID           int            `json:"roleId"`
	PermitOneAllowed bool           `json:"permitOneAllowed"`
	PermitTwoAllowed bool           `json:"permitTwoAllowed"`
	RoleName         string         `json:"rolename"`
}

//Role Role
type Role struct {
	ID   sql.NullString `json:"id"`
	Name sql.NullString `json:"name"`
}
