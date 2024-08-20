package usecase

import (
	repository "coffee/project/Repository"
	"coffee/project/domain"
)

// struct for working with the user profile
type ProfileUseCase struct {
	profile *repository.UserRepository
}

// method for handling the user profile
func NewProfileUseCase(profile *repository.UserRepository) *ProfileUseCase {
	return &ProfileUseCase{
		profile: profile,
	}
}

// method for adding profile information
func (p *ProfileUseCase) AddingProfile(id string, profile *domain.Profile) error {
	user, err := p.profile.GetUserByID(id)
	if err != nil {
		return err
	}
	if profile.FirstName != "" {
		user.FirstName = profile.FirstName
	}
	if profile.LastName != "" {
		user.LastName = profile.LastName
	}
	if profile.Email != "" {
		user.Email = profile.Email
	}
	if profile.Image != "" {
		user.Photo = profile.Image
	}

	return nil
}

// method for getting user profile from the system
func (p *ProfileUseCase) GettingProfile(id string) (*domain.Profile, error) {
	user, err := p.profile.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	var profile domain.Profile
	profile.FirstName = user.FirstName
	profile.LastName = user.LastName
	profile.Email = user.Email
	profile.Image = user.Photo
	return &profile, nil
}
