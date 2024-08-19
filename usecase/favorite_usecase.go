package usecase

import (
	repository "coffee/project/Repository"
	"coffee/project/domain"
)

// struct for handling favorite coffees for the user
type FavoriteUseCase struct {
	favoriteUseCase *repository.UserRepository
}

// method for creating new favorite usecase
func NewFavoriteUseCase(repository *repository.UserRepository) *FavoriteUseCase {
	return &FavoriteUseCase{
		favoriteUseCase: repository,
	}
}

// method for getting favorite coffee by using user id
func (f *FavoriteUseCase) GetFavoritebYId(id string) ([]*domain.Coffee, error) {
	return f.favoriteUseCase.GetFavoriteByID(id)
}

// method for adding favorite of the user in to the base
func (f *FavoriteUseCase) AddFavorite(id string, favorite *domain.Coffee) error {
	user, err := f.favoriteUseCase.GetUserByID(id)
	if err != nil {
		return err
	}
	user.Favorite = append(user.Favorite, favorite)
	newUser := domain.User{}
	newUser.Id = user.Id
	newUser.Email = user.Email
	newUser.FirstName = user.FirstName
	newUser.LastName = user.LastName
	newUser.History = user.History
	newUser.Favorite = user.Favorite
	newUser.Ongoing = user.Ongoing
	newUser.Password = user.Password

	err = f.favoriteUseCase.UpdateUser(&newUser, newUser.Id.Hex())
	return err

}

// method for getting deleting the coffe from the favorite

func (f *FavoriteUseCase) DeleteFavorite(id string, coffee string) error {
	user, err := f.favoriteUseCase.GetUserByID(id)
	if err != nil {
		return err
	}
	for index, favorite := range user.Favorite {
		if favorite.ID.Hex() == coffee {
			user.Favorite = append(user.Favorite[index:], user.Favorite...)
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

	err = f.favoriteUseCase.UpdateUser(&newUser, newUser.Id.Hex())
	return err

}
