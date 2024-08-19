package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct for handling ongoing purchases
type Ongoing struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Coffee    primitive.ObjectID `json:"coffee" bson:"_coffee,omitempty"`
	User      primitive.ObjectID `json:"user" bson:"_user,omitempty"`
	Quantity  float64            `json:"quantity" bson:"quantity"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}

// interface for handling ongoing model
type OngoingRepository interface {
	InsertOngoing(ongoing *Ongoing) error
	UpdateOngoing(Ongoing *Ongoing) error
	DeleteOngoing(id string) error
	GetAllOnGoing() ([]*Ongoing, error)
}
