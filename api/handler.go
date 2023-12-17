package api

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sologubd/currency-exchange-backend/bot"
)

type HttpHandler func(c echo.Context) error

func (a *Api) HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Alive")
}
func (a *Api) TelegramWebhook(c echo.Context) error {

	slog.Debug("New message from Telegram")
	body := c.Request().Body
	jsonBody, _ := io.ReadAll(body)

	update := &bot.Update{}
	json.Unmarshal(jsonBody, &update)

	a.bot.ParseMesage(update.Message.Text)

	return c.String(http.StatusCreated, update.Message.Text)
}
