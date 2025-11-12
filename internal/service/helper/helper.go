package helper

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/entity"
)

type Helper struct {
	Bot   *tgbotapi.BotAPI
	helps map[entity.State]string
}

func NewHelper(bot *tgbotapi.BotAPI) Helper {
	return Helper{Bot: bot, helps: Helper{}.buildHelp()}
}

func (h Helper) RootHelp(update tgbotapi.Update) error {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	//msg.ReplyToMessageID = update.Message.MessageID
	msg.Text = h.helps[entity.Root]
	_, err := h.Bot.Send(msg)
	return err
}

func (h Helper) buildHelp() map[entity.State]string {
	return map[entity.State]string{
		entity.Root:          root,
		entity.PlanTotalCost: planTotalCost,
		entity.BankRates:     bankRates,
	}
}

func (h Helper) StateHelp(state entity.State) string {
	return h.helps[state]
}
