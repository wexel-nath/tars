package market

import (
	"fmt"
	"net/http"
	"time"
)

type Trade struct{
	ID            string    `json:"id"`
	MarketID      string    `json:"marketId"`
	Timestamp     time.Time `json:"timestamp"`
	Price         string    `json:"price"`
	Amount        string    `json:"amount"`
	Side          string    `json:"side"`
	Fee           string    `json:"fee"`
	OrderID       string    `json:"orderId"`
	LiquidityType string    `json:"liquidityType"`
	ClientOrderID *string   `json:"clientOrderId"`
}

func GetTradesByOrderID(orderID string) ([]Trade, error) {
	path := "/v3/trades"

	params := map[string]string{
		"orderId": orderID,
	}

	response, err := get(path, params, nil)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request[%s] expected 200. got %d", path, response.StatusCode)
	}

	var trades []Trade
	err = unmarshalBody(response, &trades)
	return trades, err
}
