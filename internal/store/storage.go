package store

import (
	"context"
	"database/sql"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
	}
	// Comments() CommentRepository
	Users interface {
		Create(context.Context, *User) error
	}
	// Followers() FollowerRepository
	// AuthTokens() AuthTokenRepository
	// Roles() RoleRepository
}

func NewPostgresStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostgresPostStore{db},
		Users: &PostgresUserStore{db},
	}
}
