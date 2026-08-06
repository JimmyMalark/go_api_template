package httpx

import (
	"errors"
	"net/http"
	"time"

	"github.com/jimmymalark/go_api_template/internal/config"
)

func SessionToken(r *http.Request, cfg config.CookieConfig) (string, error) {
	c, err := r.Cookie(cfg.Name)
	if err != nil {
		return "", errors.New("session cookie not found")
	}

	return c.Value, nil
}

func SetSessionCookie(
	w http.ResponseWriter,
	cfg config.CookieConfig,
	token string,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     cfg.Name,
		Value:    token,
		Path:     "/",
		Domain:   cfg.Domain,
		Expires:  time.Now().Add(time.Duration(cfg.MaxAge) * time.Second),
		MaxAge:   cfg.MaxAge,
		HttpOnly: cfg.HTTPOnly,
		Secure:   cfg.Secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func ClearSessionCookie(
	w http.ResponseWriter,
	cfg config.CookieConfig,
) {
	http.SetCookie(w, &http.Cookie{
		Name:     cfg.Name,
		Value:    "",
		Path:     "/",
		Domain:   cfg.Domain,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: cfg.HTTPOnly,
		Secure:   cfg.Secure,
		SameSite: http.SameSiteLaxMode,
	})
}

func ReadSessionCookie(
	r *http.Request,
	cfg config.CookieConfig,
) (string, error) {
	cookie, err := r.Cookie(cfg.Name)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			return "", err
		}
		return "", err
	}

	return cookie.Value, nil
}
