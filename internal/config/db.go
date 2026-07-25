package config

import "fmt"

type DBConfig struct {
	Host     string `env:"DB_HOST" env-required:"true" env-description:"PostgreSQL server hostname"`
	Port     int    `env:"DB_PORT" env-default:"5432"`
	Name     string `env:"DB_NAME" env-required:"true"`
	User     string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	SSLMode  string `env:"DB_SSLMODE" env-default:"disable"`
}

func (d DBConfig) Addr() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.Name,
		d.SSLMode,
	)
}
