package config

type CookieConfig struct {
	Name     string `env:"COOKIE_NAME" env-default:"session"`
	Domain   string `env:"COOKIE_DOMAIN"`
	Secure   bool   `env:"COOKIE_SECURE" env-default:"false"`
	HTTPOnly bool   `env:"COOKIE_HTTP_ONLY" env-default:"true"`
	MaxAge   int    `env:"COOKIE_MAX_AGE" env-default:"604800"` // 7 days
}
