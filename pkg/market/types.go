package market

const (
	StatusAccepted           = "Accepted"
	StatusPlaced             = "Placed"
	StatusPartiallyMatched   = "Partially Matched"
	StatusFullyMatched       = "Fully Matched"
	StatusCancelled          = "Cancelled"
	StatusPartiallyCancelled = "Partially Cancelled"
	StatusFailed             = "Failed"

	TypeLimit      = "Limit"
	TypeMarket     = "Market"
	TypeStopLimit  = "Stop Limit"
	TypeStop       = "Stop"
	TypeTakeProfit = "Take Profit"

	SideBid = "Bid"
	SideAsk = "Ask"

	TradeLiquidityTypeMaker = "Maker"
	TradeLiquidityTypeTaker = "Taker"
)
