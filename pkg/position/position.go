package position

import (
	"fmt"
	"time"

	"tars/pkg/helper/parse"
	"tars/pkg/market"
)

type Position struct{
	ID              int64
	MarketID        string
	Price           float64
	Amount          float64
	Type            string
	Status          string
	Created         time.Time
	Updated         time.Time
	OpenFee         *float64
	CloseFee        *float64
	ExternalOrderID *string
}

func (p Position) toOrderRequest() market.OrderRequest {
	request := market.OrderRequest{
		MarketID: p.MarketID,
		Price:    fmt.Sprintf("%f", p.Price),
		Amount:   fmt.Sprintf("%f", p.Amount),
		Side:     getOrderSide(p.Type),
		Type:     market.TypeLimit, // always limit for now
	}

	return request
}

func positionFromRow(row map[string]interface{}) (p Position, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building position from row[%v]", r, row)
		}
	}()

	position := Position{
		ID:              row[columnPositionID].(int64),
		MarketID:        row[columnMarketID].(string),
		Price:           parse.MustParseFloat(row[columnPositionPrice]),
		Amount:          parse.MustParseFloat(row[columnPositionAmount]),
		Type:            row[columnPositionType].(string),
		Status:          row[columnPositionStatus].(string),
		Created:         row[columnPositionCreated].(time.Time),
		Updated:         row[columnPositionUpdated].(time.Time),
		OpenFee:         parse.FloatPointer(row[columnOpenFee]),
		CloseFee:        parse.FloatPointer(row[columnCloseFee]),
		ExternalOrderID: parse.StringPointer(row[columnExternalOrderID]),
	}

	return position, nil
}

func newPosition(
	marketID string,
	price float64,
	amount float64,
	positionType string,
	timestamp time.Time,
) (Position, error) {
	row, err := insertPosition(
		marketID,
		price,
		amount,
		positionType,
		StatusCreated,
		timestamp,
	)
	if err != nil {
		return Position{}, err
	}

	return positionFromRow(row)
}

func updatePosition(
	position Position,
	status *string,
	openFee *float64,
	closeFee *float64,
	externalOrderID *string,
) (Position, error) {
	fieldsToUpdate := make(map[string]interface{})

	if status != nil {
		fieldsToUpdate[columnPositionStatus] = *status
	}

	if openFee != nil {
		fieldsToUpdate[columnOpenFee] = *openFee
	}

	if closeFee != nil {
		fieldsToUpdate[columnCloseFee] = *closeFee
	}

	if externalOrderID != nil {
		fieldsToUpdate[columnExternalOrderID] = *externalOrderID
	}

	if len(fieldsToUpdate) == 0 {
		return position, nil
	}

	row, err := updatePositionRow(position.ID, fieldsToUpdate)
	if err != nil {
		return Position{}, err
	}

	return positionFromRow(row)
}

func Open(
	marketID string,
	price float64,
	amount float64,
	positionType string,
	timestamp time.Time,
) (Position, error) {
	p, err := newPosition(marketID, price, amount, positionType, timestamp)
	if err != nil {
		return Position{}, err
	}

	request := p.toOrderRequest()
	order, err := market.PlaceOrder(request, p.Created)
	if err != nil {
		return Position{}, err
	}

	if order.Status != market.StatusFullyMatched {
		// log/alert

		// try cancel order

		// mark position as cancelled
		status := StatusCancelled
		return updatePosition(p, &status, nil, nil, &order.ID)
	}

	orderTrades, err := market.GetTradesByOrderID(order.ID)
	if err != nil {
		return Position{}, err
	}

	totalFee := 0.0
	for _, trade := range orderTrades {
		fee, err := parse.StringToFloat(trade.Fee)
		if err != nil {
			return Position{}, err
		}

		totalFee += fee
	}

	status := StatusOpen
	return updatePosition(p, &status, &totalFee, nil, &order.ID)
}
