package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/zamyatin-zkex/estate_calc_bot/config"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/entity"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/repository"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/service/banker"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/service/denominator"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/service/helper"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/service/processor"
	"github.com/zamyatin-zkex/estate_calc_bot/internal/service/state"
	"log"
)

func main() {

	cfg := config.Build()
	bot := must(tgbotapi.NewBotAPI(cfg.Bot.Token))
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	denominatorService := denominator.NewDenominator(cfg.Bank, bot)
	help := helper.NewHelper(bot)
	bank := banker.NewBanker(cfg.Bank, bot)

	router := processor.NewRouter().
		Route(entity.Root, help.RootHelp).
		Route(entity.PlanTotalCost, denominatorService.Handle).
		Route(entity.BankRates, bank.GetRates)

	store := repository.NewState()
	proc := processor.NewProcessor(store, state.NewMachine(), bot, router, help)

	for update := range updates {
		if update.Message != nil { // If we got a message
			update := update
			go func() {
				err := proc.Process(update)
				if err != nil {
					log.Println(err)
				}
			}()
		}
	}
}
