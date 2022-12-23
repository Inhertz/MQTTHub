package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Adapter struct {
	bot    *tgbotapi.BotAPI
	target int64
}

// NewAdapter creates a new Adapter
func NewAdapter(token string, target int64) (*Adapter, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("bot connection failed: %v", err)
		return nil, err
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	return &Adapter{bot: bot, target: target}, nil
}

func (bota Adapter) PushMessage(msg string) error {
	message := tgbotapi.NewMessage(bota.target, msg)

	if _, err := bota.bot.Send(message); err != nil {
		return err
	}

	return nil
}
