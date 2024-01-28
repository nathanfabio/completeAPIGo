package userservice

import "github.com/nathanfabio/completeAPIGo/internal/repository/userepository"

func NewUserService(repo userepository.UserRepository) UserService {
	return &service{
		repo,
	}
}

type service struct {
	repo userepository.UserRepository
}

type UserService interface {
	CreateUser() error
}