package config

import (
	"context"
	"time"

	"github.com/sethvargo/go-envconfig"
)

var (
	instance Config
)

type Config struct{
	LogLevel int    `env:"LOG_LEVEL,default=20"`
	Version  string `env:"VERSION"`

	// Bot
	MarketID      string  `env:"MARKET_ID"`
	MarketBaseURL string  `env:"MARKET_BASE_URL"`
	InitialDate   string  `env:"INITIAL_DATE"`
	TickerDelta   int     `env:"TICKER_DELTA"`
	PositionSize  float64 `env:"POSITION_SIZE"`

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

func (c Config) GetTickerDelta() time.Duration {
	return time.Duration(c.TickerDelta) * time.Minute
}
