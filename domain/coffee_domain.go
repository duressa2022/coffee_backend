package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

// A model for repesenting coffee entity
type Coffee struct {
	ID          primitive.ObjectID `json:"id" bson:"_id ,omitempty"`
	Name        string             `json:"name"`
	Price       float64            `json:"price"`
	Description string             `json:"description"`
	Category    string             `json:"category"`
	Quantity    int                `json:"quantity"`
	WithMilk    bool               `json:"with_milk"`
	Rating      float64            `json:"rating"`
	Image       string             `json:"image_url"`
}

//An interface containg methods implemented by Coffee
type CoffeRepository interface {
	InsertCoffee(coffee *Coffee) error
	GetCoffeeByID(id string) (*Coffee, error)
	GetCoffeeByCategory(category string) ([]*Coffee, error)
	GetCoffeeByRating(rating float64) ([]*Coffee, error)
	GetCoffeByPrice(price float64) ([]*Coffee, error)
	UpdateCoffee(coffee *Coffee) error
	DeleteCoffeeById(id string) error
	GetName(id string) (string error)
	GetPrice(id string) (float64, error)
	GetDescription(id string) (string error)
	GetCategory(id string) (string error)
	GetQuantity(id string) (int, error)
	GetWithMilk(id string) (bool, error)
	GetRating(id string) (float64, error)
	GetImage(id string) (string, error)
}
