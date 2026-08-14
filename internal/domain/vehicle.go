package domain

import "time"

type Vehicle struct {
	ID        int    `json:"id"`
	VIN       string `json:"vin"`
	Model     string `json:"model"`
	Year      int    `json:"year"`
	CreatedAt time.Time `json:"created_at"`
}