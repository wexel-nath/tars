package position

import "time"

type Position struct{
	ID              int64
	MarketID        string
	Price           float64
	Amount          float64
	Type            string
	Status          string
	Created         time.Time
	Updated         time.Time
	ExternalOrderID string
}

func Open(
	marketID string,
	price float64,
	amount float64,
	positionType string,
) (Position, error) {
	// insert position with status "Created"

	// create order with fake-btc-markets

	// assert status is "Fully Matched"
	// update position: status, external order id, updated

	// return position
	return Position{}, nil
}
