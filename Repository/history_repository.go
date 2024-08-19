package repository

import (
	"coffee/project/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// struct for working with histroy respository
type HistoryRepository struct {
	histroyRepository *mongo.Collection
}

// method for creating new histroy repository
func NewHistoryRepository(client *mongo.Client, database string, collection string) *HistoryRepository {
	histroyRepository := client.Database(database).Collection(collection)
	return &HistoryRepository{
		histroyRepository: histroyRepository,
	}
}

// method for adding new histroy into the database
func (h *HistoryRepository) InsertHistroy(history *domain.Histroy) error {
	_, err := h.histroyRepository.InsertOne(context.TODO(), history)
	return err
}

// method for updating current histroy based id
func (u *HistoryRepository) UpdateHistroy(histroy *domain.Histroy, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	updated := bson.M{
		"quantity": histroy.PurchasedQuantity,
	}
	_, err = u.histroyRepository.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}, updated)
	return err

}

// method for deleting histroy by using history
func (u *HistoryRepository) DeleteHistroyByID(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = u.histroyRepository.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: _id}})
	return err
}

// method for deleting histroy by using coffeid
func (u *HistoryRepository) DeleteHistroyByCoffeeId(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = u.histroyRepository.DeleteOne(context.TODO(), bson.D{{Key: "_coffee_id", Value: _id}})
	return err
}

// method for deleting histroy by using userId
func (u *HistoryRepository) DeleteHistroyByUserID(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = u.histroyRepository.DeleteOne(context.TODO(), bson.D{{Key: "_user_id", Value: _id}})
	return err
}

// Method for getting histroy by using histroy id
func (u *HistoryRepository) GetHistroyByID(id string) ([]*domain.Histroy, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	cursor, err := u.histroyRepository.Find(context.TODO(), bson.D{{Key: "_id", Value: _id}})
	if err != nil {
		return nil, err
	}
	var historys []*domain.Histroy
	for cursor.Next(context.TODO()) {
		var histroy domain.Histroy
		err := cursor.Decode(&histroy)
		if err != nil {
			return nil, err
		}
		historys = append(historys, &histroy)
	}
	return historys, nil

}

// method for getting user history by using coffee id
func (u *HistoryRepository) GetHistoryByCoffeeID(id string) ([]*domain.Histroy, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	cursor, err := u.histroyRepository.Find(context.TODO(), bson.D{{Key: "_coffee_id", Value: _id}})
	if err != nil {
		return nil, err
	}
	var historys []*domain.Histroy
	for cursor.Next(context.TODO()) {
		var histroy domain.Histroy
		err := cursor.Decode(&histroy)
		if err != nil {
			return nil, err
		}
		historys = append(historys, &histroy)
	}
	return historys, nil

}

// method for getting history by using user id
func (u *HistoryRepository) GetHistoryByUserID(id string) ([]*domain.Histroy, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	cursor, err := u.histroyRepository.Find(context.TODO(), bson.D{{Key: "_user_id", Value: _id}})
	if err != nil {
		return nil, err
	}
	var historys []*domain.Histroy
	for cursor.Next(context.TODO()) {
		var histroy domain.Histroy
		err := cursor.Decode(&histroy)
		if err != nil {
			return nil, err
		}
		historys = append(historys, &histroy)
	}
	return historys, nil
}
