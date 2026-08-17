package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/srivishnu-sivan/fleet-service/internal/domain"
)

type VehicleRepository struct {
	db *pgxpool.Pool
}

func NewVehicleRepository(db *pgxpool.Pool) *VehicleRepository {
	return &VehicleRepository{
		db: db,
	}
}

func (v *VehicleRepository) CreateVehicle(ctx context.Context, vehicle domain.Vehicle) (uuid.UUID, error) {
	query := `INSERT INTO vehicle (vin, model, year)
          VALUES ($1, $2, $3) RETURNING id`

	var id uuid.UUID

	err := v.db.QueryRow(ctx, query, vehicle.VIN, vehicle.Model, vehicle.Year).Scan(&id)

	if err != nil {
		return uuid.Nil,err
	}
	return id,nil
}
