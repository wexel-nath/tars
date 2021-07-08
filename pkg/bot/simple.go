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

type simple struct{
	cycleTime         time.Duration
	previousTimestamp time.Time
	lastPurchasePrice float64
}

func NewSimple() Cycle {
	return &simple{
		cycleTime:         0,
		lastPurchasePrice: 0,
	}
}

func (s *simple) run(runID int64, timestamp time.Time) (bool, error) {
	cfg := config.Get()

	ticker, err := market.GetTickerForTimestamp(cfg.MarketID, timestamp)
	if err != nil {
		return true, err
	}

	// end of tickers
	if ticker.Timestamp == s.previousTimestamp {
		return true, nil
	}
	s.previousTimestamp = ticker.Timestamp

	// get open positions
	openPositions, err := position.GetOpenPositions(runID)
	if err != nil {
		return true, err
	}

	noOpen := len(openPositions) == 0

	price := parse.MustGetFloat(ticker.LastPrice)
	if s.lastPurchasePrice == 0.0 ||
		(noOpen && price > s.lastPurchasePrice) {
		log.Info("setting last price %f", price)
		s.lastPurchasePrice = price
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
	softEnterPrice := s.lastPurchasePrice * cfg.PositionSoftEnter
	hardEnterPrice := s.lastPurchasePrice * cfg.PositionHardEnter
	if totalExposure < cfg.MaxExposure && (
		price <= hardEnterPrice ||
		(price <= softEnterPrice && noOpen)) {

		err = placeOrder(runID, ticker)
		if err != nil {
			return true, err
		}

		s.lastPurchasePrice = price
	} else {
		log.Info(
			"not buying. price[%f] softEnterPrice[%f] hardEnterPrice[%f] openPositions[%d]",
			price,
			softEnterPrice,
			hardEnterPrice,
			len(openPositions),
		)
	}

	return false, nil
}

func placeOrder(runID int64, ticker market.Ticker) error {
	log.Info("ticker price: %s", ticker.LastPrice)

	price, err := strconv.ParseFloat(ticker.LastPrice, 64)
	if err != nil {
		return err
	}

	amount := config.Get().PositionSize / price

	log.Info("placing order. market[%s] price[%f] amount[%f]", ticker.MarketID, price, amount)

	p, err := position.Open(runID, ticker.MarketID, price, amount, position.TypeLong, ticker.Timestamp)
	if err != nil {
		return err
	}

	log.Info("position is open: %#v", p)

	return nil
}
