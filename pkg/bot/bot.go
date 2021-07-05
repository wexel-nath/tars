package bot

import (
	"fmt"
	"tars/pkg/market"
	"time"
)

type cycle interface{
	run() error
}

func Start(c cycle) error {
	//for true {
	//	err := c.run()
	//	if err != nil {
	//		return err
	//	}
	//}

	timestamp, err := time.Parse(time.RFC3339, market.InitialDate)
	if err != nil {
		return err
	}

	ticker, err := market.GetTickerForTimestamp(market.EthUSD, timestamp)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("%#v", ticker))
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
