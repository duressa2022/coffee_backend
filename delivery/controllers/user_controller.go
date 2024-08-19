package controllers

import (
	"coffee/project/domain"
	"coffee/project/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct for working with user usecase
type UserController struct {
	userUsecase *usecase.RegisterUseCase
}

// method for creating new registration
func NewRegisterController(register *usecase.RegisterUseCase) *UserController {
	return &UserController{
		userUsecase: register,
	}
}

// method for registering user in to the system
func (r *UserController) RegisterUser(c *gin.Context) {
	var userInformation domain.RegisterInfo
	if err := c.BindJSON(&userInformation); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while registering"})
		return
	}
	err := r.userUsecase.RegisterUser(&userInformation)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "error of duplication of user"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "user is succefully created"})
}

// method for deleting user from the system
func (r *UserController) DeleteUser(c *gin.Context) {
	var deleteInfo *domain.DeleteInfo
	if err := c.BindJSON(&deleteInfo); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "error while deleting the user"})
		return
	}
	err := r.userUsecase.DeleteUser(deleteInfo.Id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "incorrect delete information"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user successfully seleted from the system"})
}
