package banker

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shopspring/decimal"
	"github.com/zamyatin-zkex/estate_calc_bot/config"
	"sort"
	"strings"
)

type Banker struct {
	cfg config.Bank
	bot *tgbotapi.BotAPI
}

func NewBanker(cfg config.Bank, bot *tgbotapi.BotAPI) Banker {
	return Banker{cfg: cfg, bot: bot}
}

func (b Banker) GetRates(update tgbotapi.Update) error {

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	type year struct {
		year int
		rate decimal.Decimal
	}

	rates := make([]year, 0)

	for y, r := range b.cfg.Plan {
		rates = append(rates, year{y, r})
	}
	sort.Slice(rates, func(i, j int) bool {
		return rates[i].year < rates[j].year
	})

	buf := strings.Builder{}
	for _, rate := range rates {
		buf.WriteString(fmt.Sprintf("%d: %s\n", rate.year, rate.rate.String()))
	}

	msg.Text = buf.String()
	b.bot.Send(msg)
	return nil
}
