package bot

import (
	"strconv"
	"time"

	"tars/pkg/config"
	"tars/pkg/helper/parse"
	"tars/pkg/log"
	"tars/pkg/market"
	"tars/pkg/position"
)

type Simple struct{
	CycleTime         time.Duration
	LastPurchasePrice float64
	Market            string
}

func (s Simple) shouldBuy(price float64, totalExposure float64) bool {
	cfg := config.Get()
	return price <= s.LastPurchasePrice * cfg.PositionEnter &&
		totalExposure < cfg.MaxExposure
}

func NewSimple() *Simple {
	log.Info("%f", config.Get().PositionEnter)

	return &Simple{
		CycleTime:         0,
		LastPurchasePrice: 0,
		Market:            config.Get().MarketID,
	}
}

func (s *Simple) run(timestamp time.Time) (bool, error) {
	ticker, err := market.GetTickerForTimestamp(s.Market, timestamp)
	if err != nil {
		return true, err
	}

	price := parse.MustGetFloat(ticker.LastPrice)
	if s.LastPurchasePrice == 0.0 {
		log.Info("setting last price %f", price)
		s.LastPurchasePrice = price
	}

	// get open positions
	openPositions, err := position.GetOpenPositions()
	if err != nil {
		return true, err
	}

	totalExposure := 0.0

	// maybe trigger sells
	for _, p := range openPositions {
		if price >= p.TargetPrice() {
			log.Info("selling position[%d] at price[%f]", p.ID, price)
			_, err = p.Close(price, timestamp)
			if err != nil {
				log.Error(err)
			}
		}

		totalExposure += p.Cost()
	}

	// maybe trigger buys
	cfg := config.Get()
	enterPrice := s.LastPurchasePrice * cfg.PositionEnter
	if price <= enterPrice && totalExposure < cfg.MaxExposure {
		err = placeOrder(ticker)
		if err != nil {
			return true, err
		}

		s.LastPurchasePrice = price
	} else {
		log.Info("not buying. price[%f] enterPrice[%f]", price, enterPrice)
	}

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
