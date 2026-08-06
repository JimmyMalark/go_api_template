package apperrors

import "errors"

var (
	ErrNotFound = errors.New("not found")

	ErrEmailExists    = errors.New("email already exists")
	ErrUsernameExists = errors.New("username already exists")

	ErrForeignKey = errors.New("foreign key violation")
	ErrCheck      = errors.New("check constraint violation")

	ErrInvalidInput = errors.New("invalid input")
)

func NewUnauthorizedError(message string) error {
	return &UnauthorizedError{Message: message}
}

type UnauthorizedError struct {
	Message string
}

func (e *UnauthorizedError) Error() string {
	return e.Message
}
