package userepository

import (
	"database/sql"

	"github.com/nathanfabio/completeAPIGo/internal/database/sqlc"
)

func NewUserRepository(db *sql.DB, q *sqlc.Queries) UserRepository {
	return &repository{
		db,
		q,
	}
}

type repository struct {
	db      *sql.DB
	queries *sqlc.Queries
}

type UserRepository interface {
	CreateUser() error
}
