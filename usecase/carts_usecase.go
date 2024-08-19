package usecase

import (
	repository "coffee/project/Repository"
	"coffee/project/domain"
)

// strcut for working with carts case of the user
type CartsUseCase struct {
	cartUseCase *repository.UserRepository
}

// method for creating cart usecase
func NewCartUseCase(usecase *repository.UserRepository) *CartsUseCase {
	return &CartsUseCase{
		cartUseCase: usecase,
	}
}

// method for getting all data from the carts
func (c *CartsUseCase) GetAllFromCart(id string) ([]*domain.Coffee, error) {
	user, err := c.cartUseCase.GetUserByID(id)
	return user.Carts, err
}

// method for adding data into the carts
func (c *CartsUseCase) AddIntoCart(id string, coffee *domain.Coffee) error {
	user, err := c.cartUseCase.GetUserByID(id)
	if err != nil {
		return err
	}
	user.Carts = append(user.Carts, coffee)
	newUser := domain.User{}
	newUser.Id = user.Id
	newUser.Email = user.Email
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.History = user.History
	newUser.Favorite = user.Favorite
	newUser.Ongoing = user.Ongoing
	newUser.Password = user.Password
	newUser.Carts = user.Carts
	err = c.cartUseCase.UpdateUser(&newUser, user.Id.Hex())
	return err
}

// method for deleting coffee from the cart
func (c *CartsUseCase) DeleteFromCarts(id string, coffee string) error {
	user, err := c.cartUseCase.GetUserByID(id)
	if err != nil {
		return err
	}
	for index, cart := range user.Carts {
		if cart.ID.Hex() == coffee {
			user.Carts = append(user.Carts[:index], user.Carts[index+1:]...)
			break
		}
	}
	newUser := domain.User{}
	newUser.Id = user.Id
	newUser.Email = user.Email
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.History = user.History
	newUser.Favorite = user.Favorite
	newUser.Ongoing = user.Ongoing
	newUser.Password = user.Password
	newUser.Carts = user.Carts
	err = c.cartUseCase.UpdateUser(&newUser, user.Id.Hex())
	return err
}
