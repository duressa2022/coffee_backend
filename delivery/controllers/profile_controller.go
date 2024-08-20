package controllers

import (
	"coffee/project/domain"
	"coffee/project/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// strcut for working on the profile
type ProfileController struct {
	profileController *usecase.ProfileUseCase
}

// method for creating profile controller
func NewProfileController(profile *usecase.ProfileUseCase) *ProfileController {
	return &ProfileController{
		profileController: profile,
	}
}

// method for getting user profile information
func (p *ProfileController) GetProfile(c *gin.Context) {
	id, okay := c.Get("id")
	if !okay {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of internal"})
		return
	}
	profile, err := p.profileController.GettingProfile(id.(string))
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error of internal"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": profile})

}

// method for working on updating user profile
func (p *ProfileController) UpdateProfile(c *gin.Context) {
	id, okay := c.Get("id")
	if !okay {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal message"})
		return
	}
	var profile *domain.Profile
	if err := c.BindJSON(&profile); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "internal message"})
		return
	}
	err := p.profileController.AddingProfile(id.(string), profile)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, gin.H{"message": "error of upating"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "updated"})

}
