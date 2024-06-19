// Define all the routes.
package routes

import (
	"rest-api/controllers"
	"rest-api/middleware"

	"github.com/gin-gonic/gin"
)
func UserRoutes(router *gin.Engine, controller controllers.UserController) {
	auth:= router.Group("/auth")
	{
		auth.POST("login", controller.Login)
		auth.POST("register", controller.Register)
	}
	user := router.Group("/user")
	user.Use(middleware.ValidateToken())
	{
		user.GET("/profile",controller.FetchProfile)
	}

	
}