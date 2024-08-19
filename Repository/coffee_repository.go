package repository

import (
	"coffee/project/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// struct for repsenting  coffeee repo
type CoffeRepository struct {
	coffeCollection *mongo.Collection
}

// A function to create a new coffe repository
func NewCoffeRepository(client *mongo.Client, database string, collection string) *CoffeRepository {
	coffeeCollection := client.Database(database).Collection(collection)
	return &CoffeRepository{
		coffeCollection: coffeeCollection,
	}
}

// A method for inserting into the repo
func (c *CoffeRepository) InsertCoffee(coffee *domain.Coffee) error {
	_, err := c.coffeCollection.InsertOne(context.TODO(), coffee)
	return err
}

// A method for getting coffe by using coffe_id
func (c *CoffeRepository) GetCoffeeByID(id string) (*domain.Coffee, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var coffee *domain.Coffee
	err = c.coffeCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&coffee)
	if err != nil {
		return nil, err
	}
	return coffee, nil
}

// A method getting coffe items by using category
func (u *CoffeRepository) GetCoffeeByCategory(category string) ([]*domain.Coffee, error) {
	cursor, err := u.coffeCollection.Find(context.TODO(), bson.D{{Key: "category", Value: category}})
	if err != nil {
		return []*domain.Coffee{}, err
	}
	var coffees []*domain.Coffee
	for cursor.Next(context.TODO()) {
		var coffee domain.Coffee
		err := cursor.Decode(&coffee)
		if err != nil {
			return []*domain.Coffee{}, err
		}
		coffees = append(coffees, &coffee)
	}
	return coffees, nil
}

// A method for getting coffe items by using rating
func (u *CoffeRepository) GetCoffeeByRating(rating float64) ([]*domain.Coffee, error) {
	cursor, err := u.coffeCollection.Find(context.TODO(), bson.D{{Key: "rating", Value: rating}})
	if err != nil {
		return []*domain.Coffee{}, err
	}
	var coffees []*domain.Coffee
	for cursor.Next(context.TODO()) {
		var coffee domain.Coffee
		err := cursor.Decode(&coffee)
		if err != nil {
			return []*domain.Coffee{}, err
		}
		coffees = append(coffees, &coffee)
	}
	return coffees, nil
}

// A method for getting coffe by using price of the coffee
func (u *CoffeRepository) GetCoffeByPrice(price float64) ([]*domain.Coffee, error) {
	cursor, err := u.coffeCollection.Find(context.TODO(), bson.D{{Key: "price", Value: price}})
	if err != nil {
		return []*domain.Coffee{}, err
	}
	var coffees []*domain.Coffee
	for cursor.Next(context.TODO()) {
		var coffee domain.Coffee
		err := cursor.Decode(&coffee)
		if err != nil {
			return []*domain.Coffee{}, err
		}
		coffees = append(coffees, &coffee)
	}
	return coffees, nil
}

// A method for updating coffe items by new information
func (u *CoffeRepository) UpdateCoffee(coffee *domain.Coffee, id string) error {
	Id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	updated := bson.M{
		"name":        coffee.Name,
		"price":       coffee.Price,
		"description": coffee.Description,
		"category":    coffee.Category,
		"quantity":    coffee.Quantity,
		"with_milk":   coffee.WithMilk,
		"rating":      coffee.Rating,
		"image":       coffee.Image,
	}
	filter := bson.D{{Key: "_id", Value: Id}}
	_, err = u.coffeCollection.UpdateOne(context.TODO(), filter, updated)
	return err

}

// A method for deleting coffee item by using id
func (u *CoffeRepository) DeleteCoffeeById(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = u.coffeCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: _id}})
	return err
}

// A method to get the name of the coffe by using id
func (u *CoffeRepository) GetName(id string) (string, error) {
	cofffe, err := u.GetCoffeeByID(id)
	return cofffe.Name, err

}

// A method for getting the price of coffe by using id
func (u *CoffeRepository) GetPrice(id string) (float64, error) {
	cofffe, err := u.GetCoffeeByID(id)
	return cofffe.Price, err

}

// A method for getting the description of the coffe by using id
func (u *CoffeRepository) GetDescription(id string) (string, error) {
	cofffe, err := u.GetCoffeeByID(id)
	return cofffe.Description, err

}

// A method for getting the category of the coffe by using id
func (u *CoffeRepository) GetCategory(id string) (string, error) {
	cofffe, err := u.GetCoffeeByID(id)
	return cofffe.Category, err

}

// A method getting the quantity of the cofee by using id
func (u *CoffeRepository) GetQuantity(id string) (int, error) {
	cofffe, err := u.GetCoffeeByID(id)
	return cofffe.Quantity, err
}

// Amethod getting the type of the coffes by using id
func (u *CoffeRepository) GetWithMilk(id string) (bool, error) {
	cofffe, err := u.GetCoffeeByID(id)
	return cofffe.WithMilk, err
}

// Amethod getting the rating of the coffes by using id
func (u *CoffeRepository) GetRating(id string) (float64, error) {
	cofffe, err := u.GetCoffeeByID(id)
	return cofffe.Rating, err
}

// a Method for getting the image url by using  id
func (u *CoffeRepository) GetImage(id string) (string, error) {
	cofffe, err := u.GetCoffeeByID(id)
	return cofffe.Image, err
}
