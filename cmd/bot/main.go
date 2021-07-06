package main

import (
	"tars/pkg/bot"
	"tars/pkg/config"
	"tars/pkg/database"
	"tars/pkg/log"
)

func main() {
	err := config.Configure()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Connect()
	if err != nil {
		log.Error(err)
	}

	simpleBot := bot.NewSimple()
	err = bot.Start(simpleBot)
	if err != nil {
		log.Error(err)
	}
}
