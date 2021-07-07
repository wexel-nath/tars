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

func enterOrderSide(positionType string) string {
	if positionType == TypeShort {
		return market.SideAsk
	}

	return market.SideBid
}

func exitOrderSide(positionType string) string {
	if positionType == TypeShort {
		return market.SideBid
	}

	return market.SideAsk
}
