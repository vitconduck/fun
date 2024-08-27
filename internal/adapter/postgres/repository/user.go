package repository

import (
	"context"

	"github.com/vitconduck/fun/internal/core/domain"
	"github.com/vitconduck/fun/internal/core/port"
	"github.com/vitconduck/fun/pkg/postgres"
	"github.com/vitconduck/fun/utils"
)

type UserRepository struct {
	db *postgres.DB
}

// CreateUser implements port.UserRepository.
func (u *UserRepository) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	query := u.db.QueryBuilder.Insert("users").
		Columns("name", "email", "pasword").
		Values(user.Name, user.Email, user.Password).
		Suffix("RETURNING *")

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}
	err = u.db.QueryRow(ctx, sql, args...).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {

		if errCode := u.db.ErrorCode(err); errCode == "23505" {
			return nil, utils.ErrConflictingData
		}
	}
	return user, nil
}

// DeleteUser implements port.UserRepository.
func (u *UserRepository) DeleteUser(ctx context.Context, id uint) error {
	panic("unimplemented")
}

// FindUserByEmail implements port.UserRepository.
func (u *UserRepository) FindUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	panic("unimplemented")
}

// FindUserById implements port.UserRepository.
func (u *UserRepository) FindUserById(ctx context.Context, id uint) (*domain.User, error) {
	panic("unimplemented")
}

// ListUsers implements port.UserRepository.
func (u *UserRepository) ListUsers(ctx context.Context, skip uint, limit uint) ([]domain.User, error) {
	panic("unimplemented")
}

// UpdateUser implements port.UserRepository.
func (u *UserRepository) UpdateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	panic("unimplemented")
}

func NewUserRepository(db *postgres.DB) port.UserRepository {
	return &UserRepository{db}
}
