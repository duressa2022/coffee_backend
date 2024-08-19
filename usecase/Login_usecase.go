package usecase

import (
	infrastructure "coffee/project/Infrastructure"
	repository "coffee/project/Repository"
	"coffee/project/domain"
	"errors"
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
func (l *LoginUseCase) Login(login *domain.Login) (string, string, error) {
	condition := map[string]interface{}{
		"email": login.Email,
	}
	user, _ := l.login.GetUserByCondition(condition)

	if user == nil {
		return "", "", errors.New("error of credential")
	}
	err := infrastructure.ComparePassword(user.Password, login.Password)
	if err != nil {
		return "", "", err
	}
	var loginClaims domain.UserClaims
	loginClaims.Id = user.Id.Hex()
	loginClaims.Role = user.Role

	accessToken, err := infrastructure.GenerateAccessToken(&loginClaims)
	if err != nil {
		return "", "", err
	}
	refreshToken, err := infrastructure.GenerateRefreshToken(&loginClaims)
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil

}
