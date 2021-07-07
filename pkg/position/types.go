package position

import (
	"tars/pkg/market"
)

const (
	StatusCreated   = "Created"
	StatusCancelled = "Cancelled"
	StatusFailed    = "Failed"
	StatusOpen      = "Open"
	StatusClosed    = "Closed"

	TypeLong  = "Long"
	TypeShort = "Short"
)

func getOrderSide(positionType string) string {
	if positionType == TypeShort {
		return market.SideAsk
	}

	return market.SideBid
}
