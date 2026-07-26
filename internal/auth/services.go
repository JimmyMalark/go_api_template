package auth

import (
	"context"

	"github.com/jimmymalark/go_api_template/internal/db"
	"github.com/jimmymalark/go_api_template/internal/security"
)

type Service struct {
	store *db.Store
}

func NewService(store *db.Store) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) Register(
	ctx context.Context,
	req RegisterRequest,
	userAgent string,
	ipAddress string,
) (RegisterResult, error) {
	return db.RunInTransaction(ctx, s.store, func(r *db.Repositories) (RegisterResult, error) {
		passwordHash, err := security.Hash(req.Password)
		if err != nil {
			return RegisterResult{}, err
		}

		user, err := r.Users.CreateUser(
			ctx,
			toCreateUserParams(req, passwordHash),
		)
		if err != nil {
			return RegisterResult{}, err
		}

		token, err := security.NewSessionToken()
		if err != nil {
			return RegisterResult{}, err
		}

		tokenHash := security.HashSessionToken(token)

		params, err := toCreateSessionParams(user,
			tokenHash,
			userAgent,
			ipAddress)
		if err != nil {
			return RegisterResult{}, err
		}

		_, err = r.Sessions.CreateSession(
			ctx,
			params,
		)
		if err != nil {
			return RegisterResult{}, err
		}

		return RegisterResult{
			User:         toUserResponse(user),
			SessionToken: token,
		}, nil
	})
}
