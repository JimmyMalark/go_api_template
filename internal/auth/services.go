package auth

import (
	"context"
	"log/slog"

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

		// tokenHash := security.HashSessionToken(token)

		// params, err := toCreateSessionParams(user,
		// 	tokenHash,
		// 	userAgent,
		// 	ipAddress)
		// if err != nil {
		// 	return RegisterResult{}, err
		// }

		// _, err = r.Sessions.CreateSession(
		// 	ctx,
		// 	params,
		// )
		// if err != nil {
		// 	return RegisterResult{}, err
		// }

		return RegisterResult{
			User:         toUserResponse(user),
			SessionToken: token,
		}, nil
	})
}

func (s *Service) Login(
	ctx context.Context,
	req LoginRequest,
	userAgent string,
	ipAddress string,
) (LoginResult, error) {
	return db.RunInTransaction(ctx, s.store, func(r *db.Repositories) (LoginResult, error) {
		user, err := r.Users.GetUserByEmail(ctx, req.Email)
		if err != nil {
			return LoginResult{}, err
		}
		slog.Info("login attempt",
			"email_from_request", req.Email,
			"email_from_db", user.Email,
			"password_length", len(req.Password),
			"hash_prefix", user.PasswordHash[:10],
		)

		hash, _ := security.Hash("password123")

		err = security.ComparePassword(hash, "password123")

		if err := security.ComparePassword(user.PasswordHash, req.Password); err != nil {
			return LoginResult{}, err
		}

		token, err := security.NewSessionToken()
		if err != nil {
			return LoginResult{}, err
		}

		tokenHash := security.HashSessionToken(token)

		params, err := toCreateSessionParams(
			user,
			tokenHash,
			userAgent,
			ipAddress,
		)
		if err != nil {
			return LoginResult{}, err
		}

		_, err = r.Sessions.CreateSession(ctx, params)
		if err != nil {
			return LoginResult{}, err
		}

		return LoginResult{
			User:         toUserResponse(user),
			SessionToken: token,
		}, nil
	})
}

func (s *Service) Logout(ctx context.Context, token string) error {
	tokenHash := security.HashSessionToken(token)

	_, err := db.RunInTransaction(ctx, s.store, func(r *db.Repositories) (struct{}, error) {
		if err := r.Sessions.DeleteSessionByTokenHash(ctx, tokenHash); err != nil {
			return struct{}{}, err
		}

		return struct{}{}, nil
	})

	return err
}

type userIDContextKey struct{}

func WithUserID(ctx context.Context, userID int64) context.Context {
	return context.WithValue(ctx, userIDContextKey{}, userID)
}

func UserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(userIDContextKey{}).(int64)
	return userID, ok
}

func (s *Service) Me(
	ctx context.Context,
	identity Identity,
) (UserResponse, error) {

	return db.RunInTransaction(ctx, s.store, func(r *db.Repositories) (UserResponse, error) {

		user, err := r.Users.GetUserByID(ctx, identity.UserID)
		if err != nil {
			return UserResponse{}, err
		}

		return toUserResponse(user), nil
	})
}
