package auth

import (
	"context"

	db "github.com/jimmymalark/go_api_template/internal/db/sqlc"
)

type UserRepository interface {
	CreateUser(ctx context.Context, params db.CreateUserParams) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	GetUserByUsername(ctx context.Context, username string) (db.User, error)
	GetUserByID(ctx context.Context, id string) (db.User, error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context, params db.CreateSessionParams) (db.Session, error)
	GetSessionByTokenHash(ctx context.Context, tokenHash string) (db.Session, error)
	DeleteSession(ctx context.Context, id int64) error
	DeleteSessionByTokenHash(ctx context.Context, tokenHash string) error
	DeleteSessionsByUserID(ctx context.Context, userID int64) error
}
