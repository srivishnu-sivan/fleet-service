package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/srivishnu-sivan/fleet-service/internal/domain"
	"github.com/srivishnu-sivan/fleet-service/internal/repository"
)

type VehicleService struct {
	repo *repository.VehicleRepository
}

func NewVehicleService(repo *repository.VehicleRepository) *VehicleService{
	return &VehicleService{
		repo : repo,
	}
}


func (s *VehicleService) CreateVehicle(
	ctx context.Context,
	vehicle domain.Vehicle,
) (uuid.UUID, error) {

	return s.repo.CreateVehicle(ctx, vehicle)
}