package bot

import (
	"time"

	"tars/pkg/market"
)

type Simple struct{
	Market    string
	CycleTime time.Duration
}

func NewSimple() Simple {
	return Simple{
		Market:    market.EthUSD,
		CycleTime: 0,
	}
}

func (s Simple) run() error {
	// starting timestamp

	// get ticker price

	// get open positions

	// get total exposure

	// maybe trigger buys

	// maybe trigger sells

	time.Sleep(s.CycleTime)
	return nil
}
