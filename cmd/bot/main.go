package main

import (
	"tars/pkg/bot"
	"tars/pkg/config"
	"tars/pkg/db"
	"tars/pkg/log"
)

func main() {
	err := config.Configure()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Connect()
	if err != nil {
		log.Error(err)
	}

	simpleBot := bot.NewSimpleBot()
	err = bot.Start(simpleBot)
	if err != nil {
		log.Error(err)
	}
}
