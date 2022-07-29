package routers

import (
	"github.com/DendiAnugerah/go-jwtauth/controllers"
	"github.com/DendiAnugerah/go-jwtauth/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controllers.GetUsers())
	incomingRoutes.GET("/users/:user_id", controllers.GetUser())
}
