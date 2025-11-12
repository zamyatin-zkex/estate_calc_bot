package processor

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/entity"
)

type Router struct {
	routes map[entity.State]func(tgbotapi.Update) error
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[entity.State]func(tgbotapi.Update) error),
	}
}

func (r *Router) Route(state entity.State, fn func(tgbotapi.Update) error) *Router {
	r.routes[state] = fn
	return r
}

func (r *Router) Get(state entity.State) func(tgbotapi.Update) error {
	return r.routes[state]
}
