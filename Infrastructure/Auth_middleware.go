package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// method for working in authorization of the login process
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		access_token, err := c.Cookie("access_token")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		claims, err := VerfiyToken(access_token)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set("id", claims["id"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
