package apperrors

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

func FromPostgres(err error) error {
	var pgErr *pgconn.PgError

	if !errors.As(err, &pgErr) {
		return err
	}

	switch pgErr.Code {

	case "23505": // unique_violation

		switch pgErr.ConstraintName {

		case "users_email_key":
			return ErrEmailExists

		case "users_username_key":
			return ErrUsernameExists
		}

	case "23503":
		return ErrForeignKey

	case "23514":
		return ErrCheck

	case "22P02":
		return ErrInvalidInput
	}

	return err
}
