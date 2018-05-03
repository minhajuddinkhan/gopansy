package helpers

import (
	"gopkg.in/go-playground/validator.v9"
)

//Validate Validate
func Validate(i interface{}) error {
	v := validator.New()
	return v.Struct(i)
}
