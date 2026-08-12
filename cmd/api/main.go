package main

import (
	"github.com/gin-gonic/gin"
	internal "github.com/srivishnu-sivan/fleet-service/internal/config"
)

type CreateVehicleRequest struct {
	VehicleID   string `json:"vehicle_id" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func main() {
cfg := internal.Load()
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	router.GET("/vehicles/:id", func(c *gin.Context) {

		c.JSON(200, gin.H{
			"status":     "ok",
			"vehicle_id": c.Param("id"),
		})
	})

	router.GET("/vehicle", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":         "ok",
			"vehicle_id":     c.Query("id"),
			"page_no":        c.Query("page"),
			"vehicle_status": c.Query("status"),
			"limit":          c.Query("limit"),
		})
	})

	router.POST("/vehicles", func(c *gin.Context) {

		var req CreateVehicleRequest

		if err := c.ShouldBind(&req); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"vehicle_id":  req.VehicleID,
			"description": req.Description,
		})
	})

	router.Run(":" + cfg.Port)

}
