package config

import "fmt"

type Cache struct {
	Host string `env:"CACHE_HOST" env-default:"localhost"`
	Port int    `env:"CACHE_PORT" env-default:"5432"`
}

func (d Cache) CSN() string {
	return fmt.Sprintf(
		"%s:%d",
		d.Host,
		d.Port,
	)
}
