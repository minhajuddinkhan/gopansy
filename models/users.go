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

//UserCreateRequest UserCreateRequest
type UserCreateRequest struct {
	Username         *string `json:"username" validate:"required"`
	Email            *string `json:"email" validate:"required"`
	Password         *string `json:"password" validate:"required"`
	RoleID           *int    `json:"roleId, omitempty" validate:"required"`
	PermitOneAllowed *bool   `json:"permitOneAllowed" validate:"required"`
	PermitTwoAllowed *bool   `json:"permitTwoAllowed" validate:"required"`
}

//Role Role
type Role struct {
	ID   *string `json:"id"`
	Name *string `json:"name"`
}
