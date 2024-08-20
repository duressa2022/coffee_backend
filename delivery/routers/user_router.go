package routers

import (
	"coffee/project/delivery/controllers"

	"github.com/gin-gonic/gin"
)

// create method for setting routes used by user of apps
func SetUpRoute(router *gin.Engine, c1 *controllers.UserController,
	c2 *controllers.LoginController,
	c3 *controllers.ProfileController) {

	freeRoutees := router.Group("/")
	freeRoutees.POST("/register", c1.RegisterUser)
	freeRoutees.DELETE("/delete", c1.DeleteUser)
	freeRoutees.POST("/login", c2.LoginUser)
	freeRoutees.PUT("/user/profile/setting", c3.UpdateProfile)
	freeRoutees.GET("/user/profile", c3.UpdateProfile)

}
