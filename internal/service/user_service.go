package service

import (
	"context"

	. "go-service/internal/model"
	. "go-service/internal/repository"
)

type UserService interface {
	All(ctx context.Context) (*[]User, error)
	Load(ctx context.Context, id string) (*User, error)
	Create(ctx context.Context, user *User) (int64, error)
	Update(ctx context.Context, user *User) (int64, error)
	Patch(ctx context.Context, user map[string]interface{}) (int64, error)
	Delete(ctx context.Context, id string) (int64, error)
}

func NewUserService(repository UserRepository) UserService {
	return &userService{repository: repository}
}

type userService struct {
	repository UserRepository
}

func (s *userService) All(ctx context.Context) (*[]User, error) {
	return s.repository.All(ctx)
}

func (s *userService) Load(ctx context.Context, id string) (*User, error) {
	res, err := s.repository.Load(ctx, id)
	if err != nil {
		return nil, err
	}
	return res, err
}
func (s *userService) Create(ctx context.Context, user *User) (int64, error) {
	return s.repository.Create(ctx, user)
}
func (s *userService) Update(ctx context.Context, user *User) (int64, error) {
	return s.repository.Update(ctx, user)
}
func (s *userService) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, user)
}
func (s *userService) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
