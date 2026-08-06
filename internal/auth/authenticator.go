package auth

import (
	"context"
	"time"

	"github.com/jimmymalark/go_api_template/internal/apperrors"
	"github.com/jimmymalark/go_api_template/internal/db"
	"github.com/jimmymalark/go_api_template/internal/security"
)

type Authenticator struct {
	store *db.Store
}

func NewAuthenticator(store *db.Store) *Authenticator {
	return &Authenticator{
		store: store,
	}
}

func (a *Authenticator) Authenticate(
	ctx context.Context,
	token string,
) (Identity, error) {

	tokenHash := security.HashSessionToken(token)

	return db.RunInTransaction(ctx, a.store, func(r *db.Repositories) (Identity, error) {

		session, err := r.Sessions.GetSessionByTokenHash(ctx, tokenHash)
		if err != nil {
			return Identity{}, apperrors.NewUnauthorizedError("session not valid")
		}

		if session.ExpiresAt.Before(time.Now()) {
			return Identity{}, apperrors.NewUnauthorizedError("session has expired")
		}

		return Identity{
			UserID:    session.UserID,
			SessionID: session.ID,
		}, nil
	})
}
