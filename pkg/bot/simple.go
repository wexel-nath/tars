package bot

import (
	"tars/pkg/config"
	"tars/pkg/log"
	"tars/pkg/market"
	"tars/pkg/position"
)

type SimpleBot struct{
	enterPricePercentage  float64
	targetPricePercentage float64
	lastPrice             float64
}

func NewSimpleBot() *SimpleBot {
	cfg := config.Get()
	return &SimpleBot{
		enterPricePercentage:  cfg.PositionEnter,
		targetPricePercentage: cfg.PositionTarget,
	}
}

func (s *SimpleBot) preOpen(ticker market.Ticker) {
	if ticker.LastPrice > s.lastPrice || s.lastPrice == 0.0 {
		log.Info("setting last price %f", ticker.LastPrice)
		s.lastPrice = ticker.LastPrice
	}
}

func (s *SimpleBot) shouldOpen(ticker market.Ticker) (bool, error) {
	enterPrice := s.lastPrice * s.enterPricePercentage

	shouldOpen := ticker.LastPrice <= enterPrice
	if shouldOpen {
		s.lastPrice = ticker.LastPrice
	}

	return shouldOpen, nil
}

func (s *SimpleBot) shouldClose(ticker market.Ticker, p position.Position) (bool, error) {
	targetPrice := p.OpenPrice * s.targetPricePercentage
	shouldClose := ticker.LastPrice >= targetPrice
	if shouldClose {
		s.lastPrice = ticker.LastPrice
	}

	return shouldClose, nil
}
