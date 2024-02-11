package userservice

import (
	"context"

	"github.com/nathanfabio/completeAPIGo/internal/dto"
	"github.com/nathanfabio/completeAPIGo/internal/repository/userepository"
)

func NewUserService(repo userepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, u dto.CreateUserDto) error
}