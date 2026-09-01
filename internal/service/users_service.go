package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/srivishnu-sivan/fleet-service/internal/domain"
	"github.com/srivishnu-sivan/fleet-service/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *repository.UsersRepository
}

func NewUserService(repo *repository.UsersRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user domain.RegisterUser) (uuid.UUID, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return uuid.Nil, err
	}

	user.Password = string(hashedPassword)

	return u.repo.CreateUsers(ctx, user)

}
