package userservice

import (
	"context"

	"github.com/nathanfabio/completeAPIGo/internal/dto"
)

func (s *service) CreateUser(ctx context.Context, u dto.CreateUserDto) error {
	return nil
}

func (s *service) UpdateUser(ctx context.Context, u dto.UpdateUserDto, id string) error {
	return nil
}