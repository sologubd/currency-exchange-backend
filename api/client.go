package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/sologubd/currency-exchange-backend/bot"
)

type IApi interface {
	Start(port int)
}

type Api struct {
	bot bot.IBot
}

func New(bot bot.IBot) IApi {
	client := &Api{
		bot: bot,
	}
	return client
}

func (a *Api) Start(port int) {
	e := echo.New()
	e.GET("/", a.HomePage)
	e.POST("/telegram-webhook", a.TelegramWebhook)

	portStr := fmt.Sprintf(":%d", port)
	e.Logger.Fatal(e.Start(portStr))
}
