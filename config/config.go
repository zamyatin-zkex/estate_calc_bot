package config

import (
	"github.com/shopspring/decimal"
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Bot  Bot
	Bank Bank
}
type Bot struct {
	Token string
}

type Bank struct {
	Plan     map[int]decimal.Decimal
	LastYear int
}

func Build() Config {
	return Config{
		Bot: Bot{
			Token: os.Getenv("BOT_TOKEN"),
		},
		Bank: Bank{
			Plan: map[int]decimal.Decimal{
				2024: decimal.NewFromFloat(17.5),
				2025: decimal.NewFromFloat(19.2),
				2026: decimal.NewFromFloat(14),
				2027: decimal.NewFromFloat(8),
				2028: decimal.NewFromFloat(8),

				// ??
				2029: decimal.NewFromFloat(8),
				2030: decimal.NewFromFloat(8),
				2031: decimal.NewFromFloat(8),
			},
			LastYear: 2028,
		},
	}
}
