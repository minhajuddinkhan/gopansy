package models

//User User
type User struct {
	ID               *string `json:"user_id,omitempty"`
	Username         *string `json:"username"`
	Email            *string `json:"email"`
	HashedPassword   *string `json:"hashed_password"`
	RoleID           *int    `json:"role_id, omitempty"`
	PermitOneAllowed *bool   `json:"permit_one_allowed"`
	PermitTwoAllowed *bool   `json:"permit_two_allowed"`
	RoleName         *string `json:"role_name,omitempty"`
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
	ID   *string `json:"role_id"`
	Name *string `json:"role_name"`
}
