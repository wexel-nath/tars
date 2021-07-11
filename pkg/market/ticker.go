package market

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"tars/pkg/helper/parse"
)

type TickerJSON struct{
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

type Ticker struct{
	MarketID  string
	BestBid   float64
	BestAsk   float64
	LastPrice float64
	Volume24h float64
	Price24h  float64
	Low24h    float64
	High24h   float64
	Timestamp time.Time
}

func (ticker *Ticker) UnmarshalJSON(data []byte) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while unmarshalling ticker from data[%s]", r, string(data))
		}
	}()

	var t TickerJSON
	if err := json.Unmarshal(data, &t); err != nil {
		return err
	}

	ticker.MarketID = t.MarketID
	ticker.BestBid = parse.MustGetFloat(t.BestBid)
	ticker.BestAsk = parse.MustGetFloat(t.BestAsk)
	ticker.LastPrice = parse.MustGetFloat(t.LastPrice)
	ticker.Volume24h = parse.MustGetFloat(t.Volume24h)
	ticker.Price24h = parse.MustGetFloat(t.Price24h)
	ticker.Low24h = parse.MustGetFloat(t.Low24h)
	ticker.High24h = parse.MustGetFloat(t.High24h)
	ticker.Timestamp = t.Timestamp

	return nil
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
