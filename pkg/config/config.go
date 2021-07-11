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
	FeePercentage float64 `env:"FEE_PERCENTAGE"`
	StartDate     string  `env:"START_DATE"`
	EndDate       string  `env:"END_DATE"`
	MarketBaseURL string  `env:"MARKET_BASE_URL"`
	MarketID      string  `env:"MARKET_ID"`
	PositionSize  float64 `env:"POSITION_SIZE"`
	TickerDelta   int     `env:"TICKER_DELTA"`

	// SimpleBot
	PositionEnter  float64 `env:"POSITION_ENTER"`
	PositionTarget float64 `env:"POSITION_TARGET"`

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
