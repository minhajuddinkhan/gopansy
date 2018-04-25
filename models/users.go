package models

//User User
type User struct {
	ID               string `json:"id"`
	Username         string `json:"username"`
	HashedPassword   string `json:"hashedPassword"`
	RoleID           int    `json:"roleId"`
	PermitOneAllowed bool   `json:"permitOneAllowed"`
	PermitTwoAllowed bool   `json:"permitTwoAllowed"`
}
