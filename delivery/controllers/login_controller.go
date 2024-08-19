package controllers

import (
	infrastructure "coffee/project/Infrastructure"
	"coffee/project/domain"
	"coffee/project/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct for working woth login controller
type LoginController struct {
	loginController *usecase.LoginUseCase
}

// method for creating new login ontroller
func NewController(l *usecase.LoginUseCase) *LoginController {
	return &LoginController{
		loginController: l,
	}
}

// method for working login usecase
func (l *LoginController) LoginUser(c *gin.Context) {
	var loginDomain *domain.Login
	if err := c.BindJSON(&loginDomain); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of internal"})
		return
	}
	accessToken, refreshToken, err := l.loginController.Login(loginDomain)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error of credentials"})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	infrastructure.SetAccessRefresh(c, accessToken, refreshToken)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "you are logged into the system"})

}

// method for handling refresh token for the login
func (l *LoginController) RefreshToken(c *gin.Context) {
	id, exist := c.Get("id")
	if !exist {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of intrenal"})
		return
	}
	role, exist := c.Get("role")
	if !exist {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of internal"})
		return
	}
	refresh, err := c.Cookie("refresh_token")
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error of the authority"})
		return
	}
	if _, err := infrastructure.VerfiyToken(refresh); err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "other error for authority"})
		return
	}
	var loginClaims domain.UserClaims
	loginClaims.Id = id.(string)
	loginClaims.Role = role.(string)
	token, _ := infrastructure.GenerateAccessToken(&loginClaims)

	c.SetCookie("access_token", token, 60*15, "/", "localhost", true, true)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "token refreshed"})
}
