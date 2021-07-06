package market

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Order struct{
	ID            string    `json:"orderId"`
	MarketID      string    `json:"marketId"`
	Price         string    `json:"price"`
	Amount        string    `json:"amount"`
	Side          string    `json:"side"`
	Type          string    `json:"type"`
	TriggerPrice  *string   `json:"triggerPrice"`
	TriggerAmount *string   `json:"triggerAmount"`
	TimeInForce   string    `json:"timeInForce"`
	PostOnly      bool      `json:"postOnly"`
	SelfTrade     string    `json:"selfTrade"`
	ClientOrderID *string   `json:"clientOrderId"`
	CreationTime  time.Time `json:"creationTime"`
	Status        string    `json:"status"`
	OpenAmount    string    `json:"openAmount"`
}

type OrderRequest struct{
	MarketID      string    `json:"marketId"`
	Price         string    `json:"price"`
	Amount        string    `json:"amount"`
	Side          string    `json:"side"`
	Type          string    `json:"type"`
}

func PlaceOrder(request OrderRequest, timestamp time.Time) (Order, error) {
	path := "/v3/orders"
	params := getDefaultParams(timestamp)

	body, err := json.Marshal(request)
	if err != nil {
		return Order{}, err
	}

	response, err := post(path, body, params, nil)
	if err != nil {
		return Order{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Order{}, fmt.Errorf("request[%s] expected 200. got %d", path, response.StatusCode)
	}

	var order Order
	err = unmarshalBody(response, &order)
	return order, err
}
