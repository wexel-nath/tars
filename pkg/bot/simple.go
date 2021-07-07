package bot

import (
	"strconv"
	"time"

	"tars/pkg/config"
	"tars/pkg/log"
	"tars/pkg/market"
	"tars/pkg/position"
)

type Simple struct{
	Market    string
	CycleTime time.Duration
}

func NewSimple() Simple {
	return Simple{
		Market:    config.Get().MarketID,
		CycleTime: 0,
	}
}

func (s Simple) run(timestamp time.Time) (bool, error) {
	ticker, err := market.GetTickerForTimestamp(s.Market, timestamp)
	if err != nil {
		return true, err
	}

	// get open positions

	// get total exposure

	// maybe trigger buys
	err = placeOrder(ticker)
	if err != nil {
		return true, err
	}

	// maybe trigger sells

	return true, nil
}

func placeOrder(ticker market.Ticker) error {
	log.Info("ticker price: %s", ticker.LastPrice)

	price, err := strconv.ParseFloat(ticker.LastPrice, 64)
	if err != nil {
		return err
	}

	amount := config.Get().PositionSize / price

	log.Info("placing order. market[%s] price[%f] amount[%f]", ticker.MarketID, price, amount)

	p, err := position.Open(ticker.MarketID, price, amount, position.TypeLong, ticker.Timestamp)
	if err != nil {
		return err
	}

	log.Info("position is open: %#v", p)

	return nil
}
