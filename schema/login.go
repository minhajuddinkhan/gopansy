package schema

import (
	validator "gopkg.in/go-playground/validator.v9"
)

//Login Login
type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

//Validate Validate
func (l *Login) Validate(v *validator.Validate) error {
	return v.Struct(l)

}
