// Define all the routes.
package routes

import (
	"rest-api/controllers"

	"github.com/gin-gonic/gin"
)
func UserRoutes(router *gin.Engine, controller controllers.UserController) {
	router.POST("/api/login", controller.Login)
	router.POST("/api/register", controller.Register)
	router.DELETE("/api/profile", controller.UserService.FetchProfile)
}