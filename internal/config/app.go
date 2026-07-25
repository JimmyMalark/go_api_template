package config

type App struct {
	Port     int    `env:"PORT" env-default:"5432"`
	Host     string `env:"HOST" env-default:"localhost"`
	LogLevel string `env:"LOG_LEVEL"`
	LogJSON  bool   `env:"LOG_JSON"`
}
