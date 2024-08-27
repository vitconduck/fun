package service

import (
	"context"

	"github.com/vitconduck/fun/internal/core/domain"
	"github.com/vitconduck/fun/internal/core/port"
)

type UserService interface {
	RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUser(ctx context.Context, id uint) (*domain.User, error)
	ListUsers(ctx context.Context, skip, limit uint) ([]domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
}

type UserServiceIplm struct {
	repo port.UserRepository
}

func NewUserService(repo port.UserRepository) UserService {
	return &UserServiceIplm{repo}
}

// DeleteUser implements UserService.
func (u *UserServiceIplm) DeleteUser(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// ListUsers implements UserService.
func (u *UserServiceIplm) ListUsers(ctx context.Context, skip uint, limit uint) ([]domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u *UserServiceIplm) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	panic("unimplemented")
}

// GetUser implements UserService.
func (u *UserServiceIplm) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	var user *domain.User

	user, err := u.repo.FindUserById(ctx, id)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// RegisterUser implements UserService.
func (u *UserServiceIplm) RegisterUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	panic("unimplemented")
}
