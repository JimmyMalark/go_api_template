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
	cookie config.CookieConfig,
) *Handler {
	return &Handler{
		service:   service,
		validator: validator,
		cookie:    cookie,
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
	var req LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpx.WriteError(
			w,
			http.StatusBadRequest,
			"invalid request body",
			r,
			err,
		)
		return
	}

	result, err := h.service.Login(
		r.Context(),
		req,
		r.UserAgent(),
		httpx.ClientIP(r),
	)

	if err != nil {
		httpx.WriteError(
			w,
			http.StatusUnauthorized,
			"invalid credentials",
			r,
			err,
		)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    result.SessionToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   false, // true in production
		SameSite: http.SameSiteLaxMode,
		MaxAge:   60 * 60 * 24 * 30,
	})

	httpx.WriteJSON(
		w,
		http.StatusOK,
		result.User,
	)
}

// Logout godoc
//
// @Summary Logout
// @Description Logout the current user
// @Tags auth
// @Produce json
// @Router /auth/logout [post]
func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	token, err := httpx.SessionToken(r, h.cookie)
	if err != nil {
		httpx.WriteError(w, http.StatusUnauthorized, "not authenticated", r, err)
		return
	}

	if err := h.service.Logout(r.Context(), token); err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "failed to logout", r, err)
		return
	}

	httpx.ClearSessionCookie(w, h.cookie)

	w.WriteHeader(http.StatusNoContent)
}

// Me godoc
//
// @Summary Current user
// @Description Get the currently authenticated user
// @Tags auth
// @Produce json
// @Router /auth/me [get]
func (h *Handler) Me(w http.ResponseWriter, r *http.Request) {
	identity, ok := IdentityFromContext(r.Context())
	if !ok {
		httpx.WriteError(w, http.StatusUnauthorized, "not authenticated", r, nil)
		return
	}

	resp, err := h.service.Me(r.Context(), identity)
	if err != nil {
		httpx.WriteError(w, http.StatusInternalServerError, "failed to get user", r, err)
		return
	}

	httpx.WriteJSON(w, http.StatusOK, resp)
}
