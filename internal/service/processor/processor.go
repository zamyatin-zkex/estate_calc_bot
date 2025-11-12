package processor

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/entity"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/repository"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/service/helper"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/service/state"
)

type Processor struct {
	sm     state.Machine
	store  *repository.State
	bot    *tgbotapi.BotAPI
	router *Router
	help   helper.Helper
}

func NewProcessor(store *repository.State, sm state.Machine, bot *tgbotapi.BotAPI, router *Router, help helper.Helper) *Processor {
	return &Processor{
		sm:     sm,
		store:  store,
		bot:    bot,
		router: router,
		help:   help,
	}
}

func (p *Processor) Process(update tgbotapi.Update) error {
	sender := update.Message.From.UserName
	curState := p.store.Get(sender)
	if curState.Nil() {
		curState = entity.Root
		p.store.Set(sender, curState)
	}

	if nextState := entity.State(update.Message.Text).Parse(); !nextState.Nil() {
		if curState != nextState {
			curState = nextState
			p.store.Set(sender, curState)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, nextState.String())
			msg.Text = nextState.String() + ": \n" + p.help.StateHelp(nextState)
			p.bot.Send(msg)
		}
	}
	//
	//if nextState := entity.State(update.Message.Text).Parse(); !nextState.Nil() {
	//	p.store.Set(sender, nextState)
	//	msg := tgbotapi.NewMessage(update.Message.Chat.ID, nextState.String())
	//	msg.Text = nextState.String() + ": \n" + p.help.StateHelp(nextState)
	//	p.bot.Send(msg)
	//	return nil
	//}

	route := p.router.Get(curState)
	if route == nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "internal error")
		p.bot.Send(msg)
		return fmt.Errorf("no route: %s", curState.String())
	}

	err := route(update)
	if err != nil {
		return fmt.Errorf("route err: %v", err)
	}

	return nil
}
