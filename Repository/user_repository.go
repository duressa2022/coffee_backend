package repository

import (
	"coffee/project/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// A strcu for repsenting the user Repo
type UserRepository struct {
	userCollection *mongo.Collection
}

// A method for creating new user collection
func NewUserRepository(client *mongo.Client, database string, collection string) *UserRepository {
	userCollection := client.Database(database).Collection(collection)
	return &UserRepository{
		userCollection: userCollection,
	}
}

// A method for creating new user into the base
func (u *UserRepository) InsertUser(user *domain.User) error {
	_, err := u.userCollection.InsertOne(context.TODO(), user)
	return err
}

// A method getting user by using id
func (u *UserRepository) GetUserByID(id string) (*domain.User, error) {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user domain.User
	err = u.userCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil

}

// Amethod for getting user by using email
func (u *UserRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := u.userCollection.FindOne(context.TODO(), bson.D{{Key: "email", Value: email}}).Decode(&user)
	return &user, err
}

// Amethod for updating the user by using id
func (u *UserRepository) UpdateUser(user *domain.User, id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	updated := bson.M{
		"firstname": user.FirstName,
		"lastname":  user.LastName,
		"email":     user.Email,
		"history":   user.History,
		"ongoing":   user.Ongoing,
		"photo":     user.Photo,
		"password":  user.Password,
		"favorite":  user.Favorite,
	}
	_, err = u.userCollection.UpdateOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}, updated)
	return err
}

// Method for deleting user by using id
func (u *UserRepository) DeleteUserByID(id string) error {
	_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	var user *domain.User
	err = u.userCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: _id}}).Decode(&user)
	if user != nil {
		_, err = u.userCollection.DeleteOne(context.TODO(), bson.D{{Key: "_id", Value: _id}})
		if err != nil {
			return err
		}
	}
	return err
}

// method for getting favorite by using user id
func (u *UserRepository) GetFavoriteByID(id string) ([]*domain.Coffee, error) {
	user, err := u.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return user.Favorite, nil

}

// method for getting the user by using email
func (u *UserRepository) GetUserByCondition(condition map[string]interface{}) (*domain.User, error) {
	filter := bson.M{}
	for key, value := range condition {
		filter[key] = value
	}
	var user *domain.User
	err := u.userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err == nil || user != nil {
		return user, errors.New("error of duplication")
	}
	return nil, nil
}

// Method for getting first name bu using id
func (u *UserRepository) GetFirstName(id string) (string, error) {
	user, err := u.GetUserByID(id)
	return user.FirstName, err
}

// Method for getting lastname by using id
func (u *UserRepository) GetLastName(id string) (string, error) {
	user, err := u.GetUserByID(id)
	return user.LastName, err
}

// Method for getting email by using id
func (u *UserRepository) GetEmail(id string) (string, error) {
	user, err := u.GetUserByID(id)
	return user.Email, err
}

// Method or getting user history by using id
func (u *UserRepository) GetHistory(id string) ([]primitive.ObjectID, error) {
	user, err := u.GetUserByID(id)
	return user.History, err
}

// Method for getting user ongoing by using id
func (u *UserRepository) GetOngoing(id string) ([]primitive.ObjectID, error) {
	user, err := u.GetUserByID(id)
	return user.Ongoing, err
}

// Method for getting user photo by using id
func (u *UserRepository) GetPhoto(id string) (string, error) {
	user, err := u.GetUserByID(id)
	return user.Photo, err
}

// Method for getting user password by using id
func (u *UserRepository) GetPassword(id string) (string, error) {
	user, err := u.GetUserByID(id)
	return user.Password, err
}
