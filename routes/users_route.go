package routes

import (
	"github.com/BrunoUemura/golang-simple-api/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	users := router.Group("api/v1/users")
	{
		users.GET("/", controllers.ShowUsers)
	}
}