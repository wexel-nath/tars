package run

import (
	"encoding/json"
	"fmt"
	"time"

	"tars/pkg/config"
	"tars/pkg/helper/parse"
)

type Run struct{
	ID       int64
	Config   string
	Started  time.Time
	Finished *time.Time
}

func runFromRow(row map[string]interface{}) (run Run, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building run from row[%v]", run, row)
		}
	}()

	run = Run{
		ID:       row[columnRunID].(int64),
		Config:   row[columnRunConfig].(string),
		Started:  row[columnRunStarted].(time.Time),
		Finished: parse.TimePointer(row[columnRunFinished]),
	}

	return run, nil
}

func NewRun() (Run, error) {
	cfgBytes, err := json.Marshal(config.Get())
	if err != nil {
		return Run{}, err
	}

	row, err := insertRun(string(cfgBytes))
	if err != nil {
		return Run{}, err
	}

	return runFromRow(row)
}

func UpdateRun(runID int64) (Run, error) {
	row, err := updateRunFinished(runID)
	if err != nil {
		return Run{}, err
	}

	return runFromRow(row)
}

type Outcome struct{
	RunID           int64
	MarketID        string
	AmountHeld      float64
	LastPrice       float64
	PositionsOpened int64
	PositionsClosed int64
	TotalSpend      float64
	TotalSold       float64
	TotalFees       float64
	Profit          float64
}

func outcomeFromRow(row map[string]interface{}) (o Outcome, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic[%v] while building run from row[%v]", r, row)
		}
	}()

	outcome := Outcome{
		RunID:           row[columnRunID].(int64),
		MarketID:        row[columnMarketID].(string),
		AmountHeld:      parse.MustParseFloat(row[columnAmountHeld]),
		LastPrice:       parse.MustParseFloat(row[columnLastPrice]),
		PositionsOpened: row[columnPositionsOpened].(int64),
		PositionsClosed: row[columnPositionsClosed].(int64),
		TotalSpend:      parse.MustParseFloat(row[columnTotalSpend]),
		TotalSold:       parse.MustParseFloat(row[columnTotalSold]),
		TotalFees:       parse.MustParseFloat(row[columnTotalFees]),
		Profit:          parse.MustParseFloat(row[columnProfit]),
	}

	return outcome, nil
}

func GetOutcome(runID int64) (Outcome, error) {
	row, err := selectOutcome(runID)
	if err != nil {
		return Outcome{}, err
	}

	return outcomeFromRow(row)
}
