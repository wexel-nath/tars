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
	InitialDate       string  `env:"INITIAL_DATE"`
	MarketBaseURL     string  `env:"MARKET_BASE_URL"`
	MarketID          string  `env:"MARKET_ID"`
	MaxExposure       float64 `env:"MAX_EXPOSURE"`
	PositionSoftEnter float64 `env:"POSITION_SOFT_ENTER"`
	PositionHardEnter float64 `env:"POSITION_HARD_ENTER"`
	PositionSize      float64 `env:"POSITION_SIZE"`
	PositionTarget    float64 `env:"POSITION_TARGET"`
	TickerDelta       int     `env:"TICKER_DELTA"`

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
