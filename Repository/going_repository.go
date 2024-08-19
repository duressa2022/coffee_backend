package repository

import (
	"coffee/project/domain"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// strcut for handling ongoing repos
type OngoingRepository struct {
	ongoingCollection *mongo.Collection
}

// method for creating new ongoing repos
func NewOngoingRepository(client *mongo.Client, database string, collection string) *OngoingRepository {
	ongoingcollection := client.Database(database).Collection(collection)
	return &OngoingRepository{
		ongoingCollection: ongoingcollection,
	}
}

// method for inserting ongoing into the database
func (u *OngoingRepository) InsertOngoing(ongoing *domain.Ongoing) error {
	_, err := u.ongoingCollection.InsertOne(context.TODO(), ongoing)
	return err
}

// method for updating ongoing into the databse
func (u *OngoingRepository) UpdateOngoing(Ongoing *domain.Ongoing, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	updated := bson.M{}
	_, err = u.ongoingCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}, updated)
	return err

}

// method for deleting ongoing from the base by using id
func (u *OngoingRepository) DeleteOngoing(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = u.ongoingCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: _id}})
	return err
}

// method for getting ongoing from the base by using id
func (u *OngoingRepository) GetAllOnGoing(id string) ([]*domain.Ongoing, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	cursor, err := u.ongoingCollection.Find(context.TODO(), bson.D{{Key: "_id", Value: _id}})
	if err != nil {
		return nil, err
	}
	var ongoings []*domain.Ongoing
	for cursor.Next(context.TODO()) {
		var ongoing domain.Ongoing
		err := cursor.Decode(&ongoing)
		if err != nil {
			return nil, err
		}
		ongoings = append(ongoings, &ongoing)
	}
	return ongoings, nil
}
