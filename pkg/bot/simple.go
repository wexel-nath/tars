package bot

import (
	"tars/pkg/config"
	"tars/pkg/helper/parse"
	"tars/pkg/log"
	"tars/pkg/market"
	"tars/pkg/position"
)

type SimpleBot struct{
	lastPrice float64
}

func NewSimpleBot() *SimpleBot {
	return &SimpleBot{}
}

func (s *SimpleBot) preOpen(ticker market.Ticker) {
	price := parse.MustGetFloat(ticker.LastPrice)
	if price > s.lastPrice || s.lastPrice == 0.0 {
		log.Info("setting last price %f", price)
		s.lastPrice = price
	}
}

func (s *SimpleBot) shouldOpen(ticker market.Ticker) (bool, error) {
	price := parse.MustGetFloat(ticker.LastPrice)
	hardEnterPrice := s.lastPrice * config.Get().PositionHardEnter

	shouldOpen := price <= hardEnterPrice
	if shouldOpen {
		s.lastPrice = price
	}

	return shouldOpen, nil
}

func (s *SimpleBot) shouldClose(ticker market.Ticker, p position.Position) (bool, error) {
	price := parse.MustGetFloat(ticker.LastPrice)

	shouldClose := price >= p.TargetPrice()
	if shouldClose {
		s.lastPrice = price
	}

	return shouldClose, nil
}
