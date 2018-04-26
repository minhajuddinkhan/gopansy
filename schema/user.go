package schema

type User struct {
	Username         *string `validate:"required"`
	Email            *string `validate:"required"`
	Password         *string `validate:"required"`
	PermitOneAllowed *bool   `validate:"required"`
	PermitTwoAllowed *bool   `validate:"required"`
	RoleID           *int    `validate:"required"`
}
