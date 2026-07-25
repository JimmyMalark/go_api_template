package users

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/jimmymalark/go_api_template/internal/apperrors"
	"github.com/jimmymalark/go_api_template/internal/httpx"
	"github.com/jimmymalark/go_api_template/internal/pagination"
	"github.com/jimmymalark/go_api_template/internal/validator"
)

type Handler struct {
	service   *Service
	validator *validator.Validator
}

func NewHandler(service *Service, validator *validator.Validator) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

// ListUsers godoc
//
//	@Summary		List users
//	@Description	Get all users
//	@Tags			users
//	@Produce		json
//	@Success		200	{array}		users.UserResponse
//	@Failure		500	{object}	httpx.ErrorResponse
//	@Param			page	query		int		false	"Page number"		default(1)		minimum(1)
//	@Param			limit	query		int		false	"Items per page"	default(20)		minimum(1)	maximum(100)
//
// @Router /admin/users [get]
func (h *Handler) ListUsers(w http.ResponseWriter, r *http.Request) {
	page := 1
	limit := 20

	if p := r.URL.Query().Get("page"); p != "" {
		value, err := strconv.Atoi(p)
		if err != nil || value < 1 {
			// return 400
		}

		page = value
	}

	if ps := r.URL.Query().Get("limit"); ps != "" {
		value, err := strconv.Atoi(ps)
		if err != nil || value < 1 {
			// return 400
		}

		limit = value
	}

	users, err := h.service.ListUsers(r.Context(), pagination.Params{
		Page:  page,
		Limit: limit,
	})
	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "internal server error", r, err)
		return
	}

	_ = httpx.WriteJSON(w, http.StatusOK, users)
}

// CreateUser godoc
//
//	@Summary		Create user
//	@Description	Create a new user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		users.CreateUserRequest	true	"User payload"
//	@Success		201		{object}	users.UserResponse
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//
// @Router /admin/users [post]
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req CreateUserRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "invalid request body", r, err)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		// TODO: Add debug logger
		httpx.WriteJSON(
			w,
			http.StatusBadRequest,
			httpx.ValidationErrors(err),
		)
		return
	}

	user, err := h.service.CreateUser(r.Context(), req)
	if err != nil {
		resp := apperrors.ToHTTP(err)

		httpx.WriteJSON(w, resp.Status, resp.Body)
		return
	}

	_ = httpx.WriteJSON(w, http.StatusCreated, user)
}
