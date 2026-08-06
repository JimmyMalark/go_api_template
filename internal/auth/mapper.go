package auth

import (
	"net/netip"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/jimmymalark/go_api_template/internal/db/sqlc"
	"github.com/jimmymalark/go_api_template/internal/ids"
)

func toCreateUserParams(
	req RegisterRequest,
	passwordHash string,
) db.CreateUserParams {
	return db.CreateUserParams{
		Xid:          ids.New(),
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: passwordHash,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		BirthDate:    req.BirthDate,
	}
}

func toUserResponse(user db.User) UserResponse {
	return UserResponse{
		Xid:       user.Xid,
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		BirthDate: user.BirthDate,
	}
}

func toCreateSessionParams(
	user db.User,
	tokenHash []byte,
	userAgent string,
	ipAddress string,
) (db.CreateSessionParams, error) {

	now := time.Now()

	addr, err := netip.ParseAddr(ipAddress)
	if err != nil {
		return db.CreateSessionParams{}, err
	}

	return db.CreateSessionParams{
		UserID:     user.ID,
		TokenHash:  tokenHash,
		UserAgent:  pgtype.Text{String: userAgent, Valid: true},
		IpAddress:  &addr,
		CreatedAt:  now,
		LastUsedAt: now,
		ExpiresAt:  now.Add(7 * 24 * time.Hour),
	}, nil
}
