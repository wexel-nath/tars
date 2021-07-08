package position

import (
	"fmt"
	"strings"
	"time"

	"tars/pkg/db"
)

const (
	columnPositionID      = "position_id"
	columnRunID           = "run_id"
	columnMarketID        = "market_id"
	columnPositionType    = "position_type"
	columnPositionStatus  = "position_status"
	columnPositionCreated = "position_created"
	columnAmount          = "amount"
	columnOpenPrice       = "open_price"
	columnClosePrice      = "close_price"
	columnOpenFee         = "open_fee"
	columnCloseFee        = "close_fee"
	columnOpenOrderID     = "open_order_id"
	columnCloseOrderID    = "close_order_id"
)

var (
	positionColumns = []string{
		columnPositionID,
		columnRunID,
		columnMarketID,
		columnPositionType,
		columnPositionStatus,
		columnPositionCreated,
		columnAmount,
		columnOpenPrice,
		columnClosePrice,
		columnOpenFee,
		columnCloseFee,
		columnOpenOrderID,
		columnCloseOrderID,
	}

	positionColumnsString = strings.Join(positionColumns, ", ")
)

func insertPosition(
	runID int64,
	marketID string,
	positionType string,
	status string,
	timestamp time.Time,
	amount float64,
	price float64,
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
			$7,
			DEFAULT,
			NULL,
			NULL,
			NULL,
			NULL
		)
		RETURNING ` + positionColumnsString

	params := []interface{}{
		runID,
		marketID,
		positionType,
		status,
		timestamp,
		amount,
		price,
	}

	return db.QueryRow(query, params, positionColumns)
}

func updatePositionRow(
	positionID int64,
	u db.Updater,
) (map[string]interface{}, error) {
	queryFormat := `
		UPDATE position
		SET %s
		WHERE position_id = $%d
		RETURNING %s
	`

	updateParams, setString := u.Output(1)
	params := append(updateParams, positionID)
	query := fmt.Sprintf(queryFormat, setString, len(params), positionColumnsString)

	return db.QueryRow(query, params, positionColumns)
}

func selectOpenPositions(runID int64) ([]map[string]interface{}, error) {
	query := `
		SELECT ` + positionColumnsString + `
		FROM position
		WHERE run_id = $1
		AND position_status = 'Open'
	`

	params := []interface{}{
		runID,
	}

	return db.QueryRows(query, params, positionColumns)
}
