package g_log

import (
	"fmt"

	tlBot "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	KSystemBotAPIToken = `1851059492:AAFQAoT69k89v-JN5W5CHR7p8XlPMvM1pZM`

	KServiceSystemGroupID int64 = -563600877
)

func SendNotify(chatID int64, format string, args ...interface{}) {
	bot, err := tlBot.NewBotAPI(KSystemBotAPIToken)

	if err != nil {
		V(3).WithError(err).Infof("Create new bot over NewBotAPI error: %+v", err)
	} else {
		bot.Debug = false

		msg := tlBot.NewMessage(chatID, fmt.Sprintf(format, args...))

		bot.Send(msg)
	}
}
