package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct for working with history of transaction
type Histroy struct {
	Id                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CoffeeID          primitive.ObjectID `json:"coffee_id" bson:"_coffe_id,omitempty"`
	UserID            primitive.ObjectID `json:"user_id" bson:"_user_id,omitempty"`
	PurchasedAt       time.Time          `json:"purchased_at" bson:"puchased_at"`
	PurchasedQuantity int                `json:"quantity" bson:"quantity"`
	PurchasingMethod  string             `json:"method" bson:"method"`
}

// interface for working with Histroy transaction struct
type HistoryRepository interface {
	InsertHistroy(history *Histroy) error
	UpdateHistroy(Histroy *Histroy, id string) error
	DeleteHistroyByID(id string) error
	DeleteHistroyByCoffeeId(id string) error
	DeleteHistroyByUserID(id string) error
	GetHistroyByID(id string) error
	GetHistoryByCoffeeID(id string) error
	GetHistoryByUserID(id string) error
}
