package domain

//struct for working with login domain
type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
