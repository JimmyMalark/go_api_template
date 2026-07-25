package config

import "fmt"

type CacheConfig struct {
	Host string `env:"CACHE_HOST" env-required:"true"`
	Port int    `env:"CACHE_PORT" env-default:"6379"`
}

func (d CacheConfig) Addr() string {
	return fmt.Sprintf(
		"%s:%d",
		d.Host,
		d.Port,
	)
}
