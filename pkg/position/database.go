package position

import (
	"strings"

	"tars/pkg/database"
)

const (
	columnPositionID      = "position_id"
	columnMarketID        = "market_id"
	columnPositionPrice   = "position_price"
	columnPositionAmount  = "position_amount"
	columnPositionType    = "position_type"
	columnPositionStatus  = "position_status"
	columnPositionCreated = "position_created"
	columnPositionUpdated = "position_updated"
	columnExternalOrderID = "external_order_id"
)

var (
	positionColumns = []string{
		columnPositionID,
		columnMarketID,
		columnPositionPrice,
		columnPositionAmount,
		columnPositionType,
		columnPositionStatus,
		columnPositionCreated,
		columnPositionUpdated,
		columnExternalOrderID,
	}

	positionColumnsString = strings.Join(positionColumns, ", ")
)

func insertPosition(
	marketID string,
	price float64,
	amount float64,
	positionType string,
	status string,
) (map[string]interface{}, error) {
	query := `
		INSERT INTO position (` + positionColumnsString + `
		VALUES (
			DEFAULT,
			$1,
			$2,
			$3,
			$4,
			$5,
			DEFAULT,
			DEFAULT,
			NULL
		)
		RETURNING ` + positionColumnsString

	params := []interface{}{
		marketID,
		price,
		amount,
		positionType,
		status,
	}

	return database.QueryRow(query, params, positionColumns)
}
