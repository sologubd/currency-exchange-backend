package main

import (
	"github.com/sologubd/currency-exchange-backend/api"
	"github.com/sologubd/currency-exchange-backend/bot"
)

func main() {
	bot := bot.New()
	httpClient := api.New(bot)
	httpClient.Start(8080)
}
