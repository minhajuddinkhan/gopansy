package validators

//Login Login
type Login struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}
