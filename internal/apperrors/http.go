package apperrors

import (
	"errors"
	"net/http"
)

type Response struct {
	Status int
	Body   any
}

func ToHTTP(err error) Response {

	switch {

	case errors.Is(err, ErrEmailExists):
		return Response{
			Status: http.StatusConflict,
			Body: map[string]any{
				"errors": map[string]string{
					"email": "already exists",
				},
			},
		}

	case errors.Is(err, ErrUsernameExists):
		return Response{
			Status: http.StatusConflict,
			Body: map[string]any{
				"errors": map[string]string{
					"username": "already exists",
				},
			},
		}

	case errors.Is(err, ErrNotFound):
		return Response{
			Status: http.StatusNotFound,
			Body: map[string]string{
				"error": "resource not found",
			},
		}

	default:
		return Response{
			Status: http.StatusInternalServerError,
			Body: map[string]string{
				"error": "internal server error",
			},
		}
	}
}
