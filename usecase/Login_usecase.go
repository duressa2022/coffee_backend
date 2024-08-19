package usecase

import (
	infrastructure "coffee/project/Infrastructure"
	repository "coffee/project/Repository"
)

// struct for working on the login case
type LoginUseCase struct {
	login *repository.UserRepository
}

// method for working login usecase
func NewLoginUseCase(login *repository.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		login: login,
	}
}

// method for logging into the system
func (l *LoginUseCase) Login(email string, password string) (string, error) {
	condition := map[string]interface{}{
		"email": email,
	}
	user, err := l.login.GetUserByCondition(condition)
	if err != nil {
		return "", err
	}
	err = infrastructure.ComparePassword(user.Password, password)
	if err != nil {
		return "", err
	}
	token, err := infrastructure.GenerateToken(user)
	if err != nil {
		return "", err
	}
	return token, nil

}
