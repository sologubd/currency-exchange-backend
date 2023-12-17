package bot

import "log/slog"

type IBot interface {
	ParseMesage(msg string)
}

type MessageHandler func()

func HandleUnknown() {
	slog.Info("Unknown")
}

func HandleStart() {
	slog.Info("Start")
}

type Bot struct {
	handlers map[string]MessageHandler
}

func New() *Bot {
	bot := &Bot{}
	bot.handlers = map[string]MessageHandler{
		"/start": HandleStart,
	}

	return bot
}

func (b *Bot) ParseMesage(msg string) {
	if handler, exists := b.handlers[msg]; exists {
		handler()
	} else {
		HandleUnknown()
	}
}
