package infrastructure

import (
	"coffee/project/domain"
	"errors"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

// method for creating an access token for login or autnentication
func GenerateAccessToken(user *domain.UserClaims) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
		"exp":  time.Now().Add(time.Minute * 15).Unix(),
	}
	AccesesToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := AccesesToken.SignedString([]byte(os.Getenv("KEY")))
	return tokenString, err
}

// method for creating a refresh token for re login or authentication
func GenerateRefreshToken(user *domain.UserClaims) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", err
	}
	claims := jwt.MapClaims{
		"id":   user.Id,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24 * 7).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := refreshToken.SignedString([]byte(os.Getenv("KEY")))
	return tokenString, err

}

// method for working with verification of both access and refresh token
func VerfiyToken(tokenString string) (jwt.MapClaims, error) {
	token, _ := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if err := godotenv.Load(".env"); err != nil {
			return nil, errors.New("error while loading the token")
		}
		return []byte(os.Getenv("KEY")), nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token from the claims")
	}
	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		return nil, errors.New("expired token from the claims")

	}
	return claims, nil
}

// method for setting both access and refresh token on the cookies
func SetAccessRefresh(c *gin.Context, accessToken string, refreshToken string) {
	c.SetCookie("access_token", accessToken, 60*15, "/", "localhost", true, true)
	c.SetCookie("refresh_token", refreshToken, 60*60*24*7, "/refresh", "localhost", true, true)
}
