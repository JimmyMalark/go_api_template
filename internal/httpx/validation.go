package httpx

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationErrors(err error) ValidationErrorResponse {
	errors := make(map[string]string)

	for _, fieldErr := range err.(validator.ValidationErrors) {
		field := strings.ToLower(fieldErr.Field())

		switch fieldErr.Tag() {
		case "required":
			errors[field] = "is required"

		case "email":
			errors[field] = "must be a valid email"

		case "min":
			errors[field] = "must be at least " + fieldErr.Param() + " characters"

		case "max":
			errors[field] = "must be at most " + fieldErr.Param() + " characters"

		case "notfuture":
			errors[field] = "birth_date cannot be in the future"

		default:
			errors[field] = "is invalid"
		}
	}

	return ValidationErrorResponse{
		Errors: errors,
	}
}
