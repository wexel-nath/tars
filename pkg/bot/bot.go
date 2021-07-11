package bot

import (
	"time"

	"tars/pkg/config"
	"tars/pkg/log"
	"tars/pkg/market"
	"tars/pkg/position"
	"tars/pkg/run"
)

/** General rules
 * only make orders that can be filled instantly -- for now
 * a filled buy order is an open long position
 * a filled sell order is an open short position
 * monitor open positions
 * monitor total exposure -- keep this under max exposure?
 */

/** Configurable rules
 * decisions
 ** buy and sell signals
 ** position size
 ** maximum exposure
 */

type Bot interface{
	preOpen(ticker market.Ticker)
	shouldOpen(ticker market.Ticker) (bool, error)
	shouldClose(ticker market.Ticker, p position.Position) (bool, error)
}

func Start(bot Bot) error {
	cfg := config.Get()
	log.Info("Starting bot from date: %s", cfg.StartDate)

	timestamp, err := time.Parse(time.RFC3339, cfg.StartDate)
	if err != nil {
		return err
	}

	endDate, err := time.Parse(time.RFC3339, cfg.EndDate)
	if err != nil {
		return err
	}

	r, err := run.NewRun()
	if err != nil {
		return err
	}

	for true {
		err = handle(r.ID, timestamp, bot)
		if err != nil {
			return err
		}

		timestamp = timestamp.Add(config.Get().GetTickerDelta())
		if timestamp.After(endDate) {
			break
		}
	}

	r, err = run.UpdateRun(r.ID)
	if err != nil {
		return err
	}

	d := r.Finished.Sub(r.Started)
	log.Info("Finished in %.2f seconds", d.Seconds())

	outcome, err := run.GetOutcome(r.ID)
	if err != nil {
		return err
	}

	log.Info("Outcome: %#v", outcome)

	return nil
}

func handle(runID int64, timestamp time.Time, bot Bot) error {
	cfg := config.Get()

	ticker, err := market.GetTickerForTimestamp(cfg.MarketID, timestamp)
	if err != nil {
		return err
	}

	bot.preOpen(ticker)

	openPositions, err := position.GetOpenPositions(runID)
	if err != nil {
		return err
	}

	// maybe trigger sells
	for _, p := range openPositions {
		shouldClose, err := bot.shouldClose(ticker, p)
		if err != nil {
			return err
		}
		if shouldClose {
			log.Info("selling position[%d] at price[%f]", p.ID, ticker.LastPrice)
			_, err = p.Close(ticker.LastPrice, timestamp)
			if err != nil {
				return err
			}
		}
	}

	// maybe trigger buy
	shouldOpen, err := bot.shouldOpen(ticker)
	if err != nil {
		return err
	}
	if shouldOpen {
		err = placeOrder(runID, ticker)
		if err != nil {
			return err
		}
	} else {
		log.Info("not buying. price[%f]", ticker.LastPrice)
	}

	return nil
}

func placeOrder(runID int64, ticker market.Ticker) error {
	log.Info("ticker price: %f", ticker.LastPrice)

	amount := config.Get().PositionSize / ticker.LastPrice

	log.Info("placing order. market[%s] price[%f] amount[%f]", ticker.MarketID, ticker.LastPrice, amount)

	p, err := position.Open(runID, ticker.MarketID, ticker.LastPrice, amount, position.TypeLong, ticker.Timestamp)
	if err != nil {
		return err
	}

	log.Info("position is open: %#v", p)
	log.Info("open fee: %f", *p.OpenFee)

	return nil
}
