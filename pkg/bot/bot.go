package bot

import (
	"time"

	"tars/pkg/config"
	"tars/pkg/log"
	"tars/pkg/run"
)

type Cycle interface{
	run(runID int64, timestamp time.Time) (bool, error)
	getMaxOpenPositions() int
}

func Start(c Cycle) error {
	initialDate := config.Get().InitialDate
	log.Info("Starting bot from initial date: %s", initialDate)

	timestamp, err := time.Parse(time.RFC3339, initialDate)
	if err != nil {
		return err
	}

	r, err := run.NewRun()
	if err != nil {
		return err
	}

	for true {
		exit, err := c.run(r.ID, timestamp)
		if err != nil {
			return err
		}

		if exit {
			break
		}

		timestamp = timestamp.Add(config.Get().GetTickerDelta())
	}

	r, err = run.UpdateRun(r.ID)
	if err != nil {
		return err
	}

	d := r.Finished.Sub(r.Started)
	log.Info("Finished in %.2f seconds", d.Seconds())
	log.Info("Max Open Positions: %d", c.getMaxOpenPositions())

	outcome, err := run.GetOutcome(r.ID)
	if err != nil {
		return err
	}

	log.Info("Outcome: %#v", outcome)

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
