package bot

import (
	"time"

	"tars/pkg/config"
	"tars/pkg/log"
)

type cycle interface{
	run(timestamp time.Time) (bool, error)
}

func Start(c cycle) error {
	initialDate := config.Get().InitialDate
	log.Info("Starting bot from initial date: %s", initialDate)

	timestamp, err := time.Parse(time.RFC3339, initialDate)
	if err != nil {
		return err
	}

	exit := false
	for !exit {
		exit, err = c.run(timestamp)
		if err != nil {
			return err
		}

		timestamp = timestamp.Add(config.Get().GetTickerDelta())
	}

	return nil
}

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
