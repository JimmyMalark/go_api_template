package auth

import (
	"encoding/json"
	"net/http"

	"github.com/jimmymalark/go_api_template/internal/apperrors"
	"github.com/jimmymalark/go_api_template/internal/config"
	"github.com/jimmymalark/go_api_template/internal/httpx"
	"github.com/jimmymalark/go_api_template/internal/validator"
)

type Handler struct {
	service   *Service
	validator *validator.Validator
	cookie    config.CookieConfig
}

func NewHandler(
	service *Service,
	validator *validator.Validator,
) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
	}
}

// Register godoc
//
//	@Summary		Register
//	@Description	Register a new user account
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			request	body		RegisterRequest	true	"Registration payload"
//	@Success		201		{object}	UserResponse
//	@Failure		400		{object}	httpx.ErrorResponse
//	@Failure		409		{object}	httpx.ErrorResponse
//	@Failure		500		{object}	httpx.ErrorResponse
//
// @Router /auth/register [post]
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(w, http.StatusBadRequest, "invalid request body", r, err)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		httpx.WriteJSON(
			w,
			http.StatusBadRequest,
			httpx.ValidationErrors(err),
		)
		return
	}

	result, err := h.service.Register(
		r.Context(),
		req,
		r.UserAgent(),
		"0.0.0.0",
		// httpx.ClientIP(r),
	)
	if err != nil {
		resp := apperrors.ToHTTP(err)
		httpx.WriteJSON(w, resp.Status, resp.Body)
		return
	}

	httpx.SetSessionCookie(
		w,
		h.cookie,
		result.SessionToken,
	)

	_ = httpx.WriteJSON(
		w,
		http.StatusCreated,
		result.User,
	)
}

// Login godoc
//
// @Summary Login
// @Description Authenticate a user
// @Tags auth
// @Accept json
// @Produce json
// @Router /auth/login [post]
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// Logout godoc
//
// @Summary Logout
// @Description Logout the current user
// @Tags auth
// @Produce json
// @Router /auth/logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// Me godoc
//
// @Summary Current user
// @Description Get the currently authenticated user
// @Tags auth
// @Produce json
// @Router /auth/me [get]
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
