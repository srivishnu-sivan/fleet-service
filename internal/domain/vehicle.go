package domain

import "time"

type Vehicle struct {
	ID        int    `json:"id"`
	VIN       string `json:"vin"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	CreatedAt time.Time `json:"created_at"`
}



type CreateVehicleRequest struct {
    VIN   string `json:"vin" binding:"required"`
    Model string `json:"model" binding:"required"`
    Year  int    `json:"year" binding:"required"`
}