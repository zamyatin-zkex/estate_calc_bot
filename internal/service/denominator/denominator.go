package denominator

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/shopspring/decimal"
	"github.com/zamyatin-zkex/estate_calc_bot/config"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/entity"
	"github.com/zamyatin-zkex/estate_calc_bot/pkg/fixed"
	"time"
)

type Denominator struct {
	bank config.Bank
	bot  *tgbotapi.BotAPI
}

func NewDenominator(cfg config.Bank, bot *tgbotapi.BotAPI) Denominator {
	return Denominator{
		bank: cfg,
		bot:  bot,
	}
}

func (d Denominator) CalcTotal(plan Plan) (decimal.Decimal, error) {
	total := decimal.Zero

	for _, payDay := range plan.PayDays {
		dayTotal, err := d.CalcPayDay(plan.Start, payDay)
		if err != nil {
			return decimal.Zero, fmt.Errorf("calc day total error: %v", err)
		}
		total = total.Add(dayTotal)
	}

	return total, nil
}

func (d Denominator) CalcPayDay(start time.Time, payDay PayDay) (decimal.Decimal, error) {
	total := payDay.Amount
	one := decimal.NewFromInt(1)
	d100 := decimal.NewFromInt(100)
	yDays := decimal.NewFromInt(365)

	from := start

	for year := start.Year(); year <= payDay.Day.Year(); year++ {

		restDays := yDays.Sub(decimal.NewFromInt(int64(from.YearDay()))).Add(one)
		if year == payDay.Day.Year() {
			restDays = decimal.NewFromInt(int64(payDay.Day.YearDay())).
				Sub(decimal.NewFromInt(int64(from.YearDay()))).
				Add(decimal.NewFromInt(1))
		}

		from = from.Add(time.Hour * 24 * time.Duration(restDays.IntPart()))

		bankRate, ok := d.bank.Plan[year]
		if !ok {
			return one, fmt.Errorf("rate for year %d not found", year)
		}

		rate := bankRate.Div(yDays).Mul(restDays) // procent
		rate = rate.Div(d100).Add(one)            // koef

		total = total.Div(rate)

		// x * r1 * r2 * r3 == 100
		// x = 100 / r1 / r2 ...

		// r1 ==  (rate/365) * (365-now) + 1

		// 73.277851792824835 * (1.192/7)

	}

	return total, nil
}

func (d Denominator) Handle(update tgbotapi.Update) error {
	if entity.State(update.Message.Text).Parse() == entity.PlanTotalCost {
		return nil
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
	msg.ReplyToMessageID = update.Message.MessageID

	///
	plan, err := NewPlanFromRaw(update.Message.Text)
	if err != nil {
		msg.Text = err.Error()
		_, _ = d.bot.Send(msg)
		return err
	}
	err = plan.Valid()
	if err != nil {
		msg.Text = err.Error()
		_, _ = d.bot.Send(msg)
		return err
	}

	total, err := d.CalcTotal(plan)
	if err != nil {
		msg.Text = err.Error()
		_, _ = d.bot.Send(msg)
		return err
	}
	msg.Text = fixed.Pad(total)

	_, _ = d.bot.Send(msg)
	return nil
}
