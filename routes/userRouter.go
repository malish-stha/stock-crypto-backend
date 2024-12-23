package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/malish-stha/stock-crypto-backend/controllers"
	"github.com/malish-stha/stock-crypto-backend/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("/users", controller.GetUsers())
	incomingRoutes.DELETE("/users/:user_id", controller.GetUser())
} 