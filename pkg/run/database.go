package run

import (
	"strings"

	"tars/pkg/db"
)

const (
	columnRunID       = "run_id"
	columnRunConfig   = "run_config"
	columnRunStarted  = "run_started"
	columnRunFinished = "run_finished"

	columnMarketID        = "market_id"
	columnAmountHeld      = "amount_held"
	columnLastPrice       = "last_price"
	columnPositionsOpened = "positions_opened"
	columnPositionsClosed = "positions_closed"
	columnTotalSpend      = "total_spend"
	columnTotalSold       = "total_sold"
	columnTotalFees       = "total_fees"
	columnProfit          = "profit"
)

var (
	runColumns = []string{
		columnRunID,
		columnRunConfig,
		columnRunStarted,
		columnRunFinished,
	}

	runColumnsString = strings.Join(runColumns, ", ")

	outcomeColumns = []string{
		columnRunID,
		columnMarketID,
		columnAmountHeld,
		columnLastPrice,
		columnPositionsOpened,
		columnPositionsClosed,
		columnTotalSpend,
		columnTotalSold,
		columnTotalFees,
		columnProfit,
	}
)

func insertRun(config string) (map[string]interface{}, error) {
	query := `
		INSERT INTO run (` + runColumnsString + `)
		VALUES (
			DEFAULT,
			$1,
			DEFAULT,
			NULL
		)
		RETURNING ` + runColumnsString

	params := []interface{}{
		config,
	}

	return db.QueryRow(query, params, runColumns)
}

func updateRunFinished(runID int64) (map[string]interface{}, error) {
	query := `
		UPDATE run
		SET run_finished = NOW()
		WHERE run_id = $1
		RETURNING ` + runColumnsString

	params := []interface{}{
		runID,
	}

	return db.QueryRow(query, params, runColumns)
}

func selectOutcome(runID int64) (map[string]interface{}, error) {
	query := `
		WITH run_positions AS (
			SELECT *
			FROM position
			WHERE run_id = $1
			AND position_status IN ('Open', 'Closed')
		),
		open_positions AS (
			SELECT *
			FROM run_positions
			WHERE position_status = 'Open'
		),
		final_positions AS (
			SELECT
				run_id,
				market_id,
				(SELECT SUM(amount) FROM open_positions) AS amount_held,
				COUNT(*) AS positions_opened,
				(SELECT COUNT(*) FROM run_positions WHERE position_status = 'Closed') AS positions_closed,
				SUM(open_price * amount) AS total_spend,
				SUM(close_price * amount) AS total_sold,
				SUM(open_fee + close_fee) AS total_fees
			FROM run_positions
			GROUP BY run_id, market_id
		)
		SELECT
			run_id,
			market_id,
			COALESCE(amount_held, 0),
			0.0 AS last_price,
			positions_opened,
			positions_closed,
			total_spend,
			total_sold,
			total_fees,
			total_sold - (total_spend + total_fees) AS profit
		FROM final_positions
	`

	params := []interface{}{
		runID,
	}

	return db.QueryRow(query, params, outcomeColumns)
}
