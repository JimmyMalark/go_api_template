package config

import "fmt"

type AppConfig struct {
	Port     int    `env:"PORT" env-default:"8080"`
	Host     string `env:"HOST" env-default:"localhost"`
	Env      string `env:"APP_ENV" env-default:"development"`
	LogLevel string `env:"LOG_LEVEL" env-default:"info"`
	LogJSON  bool   `env:"LOG_JSON" env-default:"false"`
}

func (c AppConfig) Validate() error {
	switch c.Env {
	case "development", "test", "production":
	default:
		return fmt.Errorf("APP_ENV must be development, test or production")
	}

	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("invalid PORT")
	}

	return nil
}
