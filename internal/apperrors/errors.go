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
