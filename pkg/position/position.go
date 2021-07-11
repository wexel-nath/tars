package position

import (
	"fmt"
	"time"

	"tars/pkg/config"
	"tars/pkg/db"
	"tars/pkg/helper/parse"
	"tars/pkg/log"
	"tars/pkg/market"
)

type Position struct{
	ID           int64
	MarketID     string
	Type         string
	Status       string
	Created      time.Time
	Amount       float64
	OpenPrice    float64
	ClosePrice   *float64
	OpenFee      *float64
	CloseFee     *float64
	OpenOrderID  *string
	CloseOrderID *string
}

func (p Position) Cost() float64 {
	return p.OpenPrice * p.Amount
}

func (p Position) enterOrderRequest() market.OrderRequest {
	request := market.OrderRequest{
		MarketID: p.MarketID,
		Price:    fmt.Sprintf("%f", p.OpenPrice),
		Amount:   fmt.Sprintf("%f", p.Amount),
		Side:     enterOrderSide(p.Type),
		Type:     market.TypeLimit, // always limit for now
	}

	return request
}

func (p Position) exitOrderRequest(price float64) market.OrderRequest {
	request := market.OrderRequest{
		MarketID: p.MarketID,
		Price:    fmt.Sprintf("%f", price),
		Amount:   fmt.Sprintf("%f", p.Amount),
		Side:     exitOrderSide(p.Type),
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
		ID:           row[columnPositionID].(int64),
		MarketID:     row[columnMarketID].(string),
		Type:         row[columnPositionType].(string),
		Status:       row[columnPositionStatus].(string),
		Created:      row[columnPositionCreated].(time.Time),
		Amount:       parse.MustParseFloat(row[columnAmount]),
		OpenPrice:    parse.MustParseFloat(row[columnOpenPrice]),
		ClosePrice:   parse.FloatPointer(row[columnClosePrice]),
		OpenFee:      parse.FloatPointer(row[columnOpenFee]),
		CloseFee:     parse.FloatPointer(row[columnCloseFee]),
		OpenOrderID:  parse.StringPointer(row[columnOpenOrderID]),
		CloseOrderID: parse.StringPointer(row[columnCloseOrderID]),
	}

	return position, nil
}

func positionsFromRows(rows []map[string]interface{}) ([]Position, error) {
	positions := make([]Position, 0)

	for _, row := range rows {
		p, err := positionFromRow(row)
		if err != nil {
			log.Error(err)
			continue
		}

		positions = append(positions, p)
	}

	return positions, nil
}

func newPosition(
	runID int64,
	marketID string,
	positionType string,
	amount float64,
	price float64,
	timestamp time.Time,
) (Position, error) {
	row, err := insertPosition(
		runID,
		marketID,
		positionType,
		StatusCreated,
		timestamp,
		amount,
		price,
	)
	if err != nil {
		return Position{}, err
	}

	return positionFromRow(row)
}

func updatePosition(position Position, u db.Updater) (Position, error) {
	if !u.ShouldUpdate() {
		return position, nil
	}

	row, err := updatePositionRow(position.ID, u)
	if err != nil {
		return Position{}, err
	}

	return positionFromRow(row)
}

func Open(
	runID int64,
	marketID string,
	price float64,
	amount float64,
	positionType string,
	timestamp time.Time,
) (Position, error) {
	p, err := newPosition(runID, marketID, positionType, amount, price, timestamp)
	if err != nil {
		return Position{}, err
	}

	//request := p.enterOrderRequest()
	//order, err := market.PlaceOrder(request, p.Created)
	//if err != nil {
	//	return Position{}, err
	//}
	//
	//if order.Status != market.StatusFullyMatched {
	//	// log/alert
	//
	//	// try cancel order
	//
	//	// mark position as cancelled
	//	u := db.NewUpdater().
	//		Set(columnPositionStatus, StatusCancelled).
	//		Set(columnOpenOrderID, order.ID)
	//	return updatePosition(p, u)
	//}
	//
	//orderTrades, err := market.GetTradesByOrderID(order.ID)
	//if err != nil {
	//	return Position{}, err
	//}
	//
	//openFee := 0.0
	//for _, trade := range orderTrades {
	//	openFee += parse.MustGetFloat(trade.Fee)
	//}

	u := db.NewUpdater().
		Set(columnPositionStatus, StatusOpen).
		Set(columnOpenFee, calculateTradeFee(price, amount)).
		Set(columnOpenOrderID, "")
	return updatePosition(p, u)
}

func (p Position) Close(price float64, timestamp time.Time) (Position, error) {
	//request := p.exitOrderRequest(price)
	//order, err := market.PlaceOrder(request, timestamp)
	//if err != nil {
	//	return Position{}, err
	//}
	//
	//if order.Status != market.StatusFullyMatched {
	//	// log/alert
	//
	//	// try cancel order | position_event??
	//
	//	// leave position open
	//	return Position{}, fmt.Errorf("failed closing position[%d]", p.ID)
	//}
	//
	//orderTrades, err := market.GetTradesByOrderID(order.ID)
	//if err != nil {
	//	return Position{}, err
	//}
	//
	//closeFee := 0.0
	//for _, trade := range orderTrades {
	//	closeFee += parse.MustGetFloat(trade.Fee)
	//}

	u := db.NewUpdater().
		Set(columnPositionStatus, StatusClosed).
		Set(columnCloseFee, calculateTradeFee(price, p.Amount)).
		Set(columnClosePrice, price).
		Set(columnCloseOrderID, "")
	return updatePosition(p, u)
}

func GetOpenPositions(runID int64) ([]Position, error) {
	rows, err := selectOpenPositions(runID)
	if err != nil {
		return nil, err
	}

	return positionsFromRows(rows)
}

func calculateTradeFee(price float64, amount float64) float64 {
	return price * amount * config.Get().FeePercentage
}
