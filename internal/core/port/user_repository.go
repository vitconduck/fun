package port

import (
	"context"

	"github.com/vitconduck/fun/internal/core/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	FindUserById(ctx context.Context, id uint) (*domain.User, error)
	FindUserByEmail(ctx context.Context, email string) (*domain.User, error)
	ListUsers(ctx context.Context, skip, limit uint) ([]domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	DeleteUser(ctx context.Context, id uint) error
}
