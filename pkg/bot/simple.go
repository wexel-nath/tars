package bot

import (
	"tars/pkg/config"
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
	if ticker.LastPrice > s.lastPrice || s.lastPrice == 0.0 {
		log.Info("setting last price %f", ticker.LastPrice)
		s.lastPrice = ticker.LastPrice
	}
}

func (s *SimpleBot) shouldOpen(ticker market.Ticker) (bool, error) {
	hardEnterPrice := s.lastPrice * config.Get().PositionHardEnter

	shouldOpen := ticker.LastPrice <= hardEnterPrice
	if shouldOpen {
		s.lastPrice = ticker.LastPrice
	}

	return shouldOpen, nil
}

func (s *SimpleBot) shouldClose(ticker market.Ticker, p position.Position) (bool, error) {
	shouldClose := ticker.LastPrice >= p.TargetPrice()
	if shouldClose {
		s.lastPrice = ticker.LastPrice
	}

	return shouldClose, nil
}
