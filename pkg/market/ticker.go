package market

import (
	"fmt"
	"net/http"
	"time"
)

type Ticker struct{
	MarketID  string    `json:"marketId"`
	BestBid   string    `json:"bestBid"`
	BestAsk   string    `json:"bestAsk"`
	LastPrice string    `json:"lastPrice"`
	Volume24h string    `json:"volume24h"`
	Price24h  string    `json:"price24h"`
	Low24h    string    `json:"low24h"`
	High24h   string    `json:"high24h"`
	Timestamp time.Time `json:"timestamp"`
}

func GetTickerForTimestamp(marketID string, timestamp time.Time) (Ticker, error) {
	path := fmt.Sprintf("/v3/markets/%s/ticker", marketID)
	params := getDefaultParams(timestamp)

	response, err := get(path, params, nil)
	if err != nil {
		return Ticker{}, err
	}

	if response.StatusCode != http.StatusOK {
		return Ticker{}, fmt.Errorf("request[%s] expected 200. got %d", path, response.StatusCode)
	}

	var ticker Ticker
	err = unmarshalBody(response, &ticker)
	return ticker, err
}
