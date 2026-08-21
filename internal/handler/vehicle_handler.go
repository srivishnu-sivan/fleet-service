package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/srivishnu-sivan/fleet-service/internal/domain"
	"github.com/srivishnu-sivan/fleet-service/internal/service"
)


type VehicleHandler struct {
	service *service.VehicleService
}

func NewVehicleHandler(service *service.VehicleService) *VehicleHandler {
	return &VehicleHandler{
		service: service,
	}
}

func (h *VehicleHandler) CreateVehicle(c *gin.Context) {
	// 1. bind JSON
	var req domain.CreateVehicleRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request fields: " + err.Error()})
		return
	}

// 2. create domain.Vehicle
	vehicle := domain.Vehicle{
    VIN:   req.VIN,
    Model: req.Model,
    Year:  req.Year,
}
	// 3. call service
	// id, err := h.service.CreateVehicle(c.Request.Context(), vehicle)

	id, err := h.service.CreateVehicle(c.Request.Context(), vehicle)
	if err != nil{
		c.JSON(500, gin.H{"error": "Server failed to save vehicle"})
		return
	}
	// 4. return response
	 c.JSON(201, gin.H{"id": id})
}
