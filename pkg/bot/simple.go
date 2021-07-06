package bot

import (
	"fmt"
	"strconv"
	"time"

	"tars/pkg/config"
	"tars/pkg/log"
	"tars/pkg/market"
)

type Simple struct{
	Market    string
	CycleTime time.Duration
}

func NewSimple() Simple {
	return Simple{
		Market:    config.Get().MarketID,
		CycleTime: 0,
	}
}

func (s Simple) run(timestamp time.Time) (bool, error) {
	ticker, err := market.GetTickerForTimestamp(s.Market, timestamp)
	if err != nil {
		return true, err
	}

	// get open positions

	// get total exposure

	// maybe trigger buys
	err = placeOrder(ticker)
	if err != nil {
		return true, err
	}

	// maybe trigger sells

	return true, nil
}

func placeOrder(ticker market.Ticker) error {
	log.Info("ticker price: %s", ticker.LastPrice)

	request, err := buildBuyOrderRequest(ticker.LastPrice)
	if err != nil {
		return err
	}

	log.Info("placing order. market[%s] amount[%s]", request.MarketID, request.Amount)

	order, err := market.PlaceOrder(request, ticker.Timestamp)
	if err != nil {
		return err
	}

	log.Info("order status: %s", order.Status)

	orderTrades, err := market.GetTradesByOrderID(order.ID)
	if err != nil {
		return err
	}

	for _, trade := range orderTrades {
		log.Info("trade fee: %s", trade.Fee)
	}

	return nil
}

func buildBuyOrderRequest(
	price string,
) (market.OrderRequest, error) {
	cfg := config.Get()

	priceFloat, err := strconv.ParseFloat(price, 64)
	if err != nil {
		return market.OrderRequest{}, err
	}

	amount := cfg.PositionSize / priceFloat

	request := market.OrderRequest{
		MarketID: cfg.MarketID,
		Price:    price,
		Amount:   fmt.Sprintf("%f", amount),
		Side:     market.SideBid,
		Type:     market.TypeLimit,
	}

	return request, nil
}
