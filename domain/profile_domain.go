package domain

//struct for working with profile domain
type Profile struct {
	FirstName string `json:"firstname" bson:"firstname"`
	LastName  string `json:"lastname" bson:"lastname"`
	Email     string `json:"email" bson:"email"`
	Image     string `json:"image" bson:"image"`
}
