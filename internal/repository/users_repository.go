package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/srivishnu-sivan/fleet-service/internal/domain"
)

type UsersRepository struct {
	db *pgxpool.Pool
}

func NewUsersRepository(db *pgxpool.Pool) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

func (u *UsersRepository) CreateUsers(ctx context.Context, users domain.RegisterUser) (uuid.UUID, error) {

	query := `INSERT INTO users(email,
password,
role_id) VALUES($1,$2,$3) RETURNING id`

	var id uuid.UUID

	err := u.db.QueryRow(ctx, query, users.Email, users.Password, users.RoleID)

	if err != nil {
		return uuid.Nil, err

	}
	return id, nil

}
