package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

var (
	instance Config
)

type Config struct{
	LogLevel int    `env:"LOG_LEVEL,default=20"`
	Version  string `env:"VERSION"`

	// Database
	DatabaseHost string `env:"DB_HOST"`
	DatabaseName string `env:"DB_NAME"`
	DatabasePass string `env:"DB_PASS"`
	DatabasePort int    `env:"DB_PORT"`
	DatabaseUser string `env:"DB_USER"`
}

func Configure() error {
	return envconfig.Process(context.Background(), &instance)
}

func Get() Config {
	return instance
}
