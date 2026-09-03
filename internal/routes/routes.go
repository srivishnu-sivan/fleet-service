package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/srivishnu-sivan/fleet-service/internal/auth"
	"github.com/srivishnu-sivan/fleet-service/internal/handler"
)

func Register(
	router *gin.Engine,
	userHandler *handler.UserHandler,
	vehicleHandler *handler.VehicleHandler,
	jwtService *auth.JWTService,
) {
	router.POST("/register", userHandler.CreateUser)

	router.POST(
		"/vehicles",
		jwtService.AuthMiddleware(),
		vehicleHandler.CreateVehicle,
	)
}