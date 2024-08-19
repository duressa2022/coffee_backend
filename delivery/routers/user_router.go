package routers

import (
	"coffee/project/delivery/controllers"

	"github.com/gin-gonic/gin"
)

// create method for setting routes used by user of apps
func SetUpRoute(router *gin.Engine, controller *controllers.UserController) {

	freeRoutees := router.Group("/")
	freeRoutees.POST("/register", controller.RegisterUser)
	freeRoutees.DELETE("/delete", controller.DeleteUser)

}
