package usecase

import (
	infrastructure "coffee/project/Infrastructure"
	repository "coffee/project/Repository"
	"coffee/project/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// struct for working with the registration of the user
type RegisterUseCase struct {
	register *repository.UserRepository
}

// method for creating new register usecase
func NewRegisterUSeCase(register *repository.UserRepository) *RegisterUseCase {
	return &RegisterUseCase{
		register: register,
	}
}

// method for adding user into the base
func (r *RegisterUseCase) RegisterUser(information *domain.RegisterInfo) error {
	condition := map[string]interface{}{
		"email": information.Email,
	}
	_, err := r.register.GetUserByCondition(condition)
	if err != nil {
		return err
	}
	var userAccount domain.User
	userAccount.Id = primitive.NewObjectID()
	hashedPassword, err := infrastructure.HashPassword(information.Password)
	if err != nil {
		return err
	}
	userAccount.FirstName = information.FirstName
	userAccount.LastName = information.LastName
	userAccount.Email = information.Email
	userAccount.Password = hashedPassword
	err = r.register.InsertUser(&userAccount)
	return err
}

// method for deleting user from the database
func (r *RegisterUseCase) DeleteUser(id string) error {
	return r.register.DeleteUserByID(id)
}

// method fro getting user based on the id
func (r *RegisterUseCase) GetUser(id string) (*domain.User, error) {
	return r.register.GetUserByID(id)
}
