package models

//User User
type User struct {
	ID               *string `json:"id,omitempty"`
	Username         *string `json:"username"`
	Email            *string `json:"email"`
	HashedPassword   *string `json:"hashedPassword"`
	RoleID           *int    `json:"roleId, omitempty"`
	PermitOneAllowed *bool   `json:"permitOneAllowed"`
	PermitTwoAllowed *bool   `json:"permitTwoAllowed"`
	RoleName         *string `json:"rolename,omitempty"`
}

type UserCreateRequest struct {
	Username         *string `json:"username"`
	Email            *string `json:"email"`
	Password         *string `json:"password"`
	RoleID           *int    `json:"roleId, omitempty"`
	PermitOneAllowed *bool   `json:"permitOneAllowed"`
	PermitTwoAllowed *bool   `json:"permitTwoAllowed"`
}

//Role Role
type Role struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}
