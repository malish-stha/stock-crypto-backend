package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/malish-stha/stock-crypto-backend/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("users/signup", controller.Signuo())
	incomingRoutes.POST("users/login", controller.Login())
}