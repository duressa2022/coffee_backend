package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// a struct or model for repsenting the user model/entity
type User struct {
	Id        primitive.ObjectID   `json:"id" bson:"_id"`
	FirstName string               `json:"firstname" bson:"firstname"`
	LastName  string               `json:"lastname" bson:"lastname"`
	Email     string               `json:"email" bson:"email,email"`
	History   []primitive.ObjectID `json:"history" bson:"history"`
	Ongoing   []primitive.ObjectID `json:"ongoing" bson:"ongoing"`
	Photo     string               `json:"photo" bson:"photo"`
	Password  string               `json:"password" bson:"password,omitempty"`
	Favorite  []*Coffee            `json:"favorite" bson:"favorite"`
	Carts     []*Coffee            `json:"carts" bson:"carts"`
	Role      string               `json:"role" bson:"role"`
}

// interface contaning the methods for working with user
type UserRepository interface {
	InsertUser(user *User) error
	GetUserByID(is string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user *User) error
	DeleteUserByID(id string) error
	GetFirstName(id string) (string, error)
	GetLastName(id string) (string, error)
	GetEmail(id string) (string, error)
	GetHistory(id string) ([]*Coffee, error)
	GetOngoing(id string) ([]*Coffee, error)
	GetPhoto(id string) (string, error)
	GetPassword(id string) (string, error)
}
