package position

import (
	"fmt"
	"strings"
	"time"

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
	columnOpenFee         = "open_fee"
	columnCloseFee        = "close_fee"
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
		columnOpenFee,
		columnCloseFee,
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
	timestamp time.Time,
) (map[string]interface{}, error) {
	query := `
		INSERT INTO position (` + positionColumnsString + `)
		VALUES (
			DEFAULT,
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			DEFAULT,
			NULL,
			NULL,
			NULL
		)
		RETURNING ` + positionColumnsString

	params := []interface{}{
		marketID,
		price,
		amount,
		positionType,
		status,
		timestamp,
	}

	return database.QueryRow(query, params, positionColumns)
}

func updatePositionRow(
	positionID int64,
	fieldsToUpdate map[string]interface{},
) (map[string]interface{}, error) {
	params := []interface{}{
		positionID,
	}

	setParts := make([]string, 0)
	placeholder := 2
	for key, value := range fieldsToUpdate {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", key, placeholder))
		placeholder++
		params = append(params, value)
	}

	query := `
		UPDATE position
		SET ` + strings.Join(setParts, ", ") + `
		WHERE position_id = $1
		RETURNING ` + positionColumnsString

	return database.QueryRow(query, params, positionColumns)
}

func selectOpenPositions() ([]map[string]interface{}, error) {
	query := `
		SELECT ` + positionColumnsString + `
		FROM position
		WHERE position_status = 'Open'
	`

	return database.QueryRows(query, nil, positionColumns)
}
