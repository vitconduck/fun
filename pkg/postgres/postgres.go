package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/vitconduck/fun/pkg/configs"
)

type DB struct {
	*pgxpool.Pool
	QueryBuilder *squirrel.StatementBuilderType
	url          string
}

func New(ctx context.Context, cfg *configs.DB) (*DB, error) {
	url := fmt.Sprintf("%s://%s:%s@%s/%s?sslmode=disable",
		"postgres",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBAddress,
		cfg.DBName,
	)

	db, err := pgxpool.New(ctx, url)

	if err != nil {
		return nil, err

	}

	err = db.Ping(ctx)
	if err != nil {
		return nil, err
	}
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	return &DB{
		Pool:         db,
		QueryBuilder: &psql,
		url:          url,
	}, nil

}

// ErrorCode returns the error code of the given error
func (db *DB) ErrorCode(err error) string {
	pgErr := err.(*pgconn.PgError)
	return pgErr.Code
}
