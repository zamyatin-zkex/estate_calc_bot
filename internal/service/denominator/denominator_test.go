package denominator

import (
	"github.com/shopspring/decimal"
	"github.com/zamyatin-zkex/estate_calc_bot/config"
	"testing"
	"time"
)

func TestDenominator_CalcPayDay_Year(t *testing.T) {

	cfg := config.Build()
	den := NewDenominator(cfg.Bank, nil)

	payDay := PayDay{
		Day:    time.Now().AddDate(3, 0, 0),
		Amount: decimal.NewFromInt(100),
	}
	t.Log(payDay.Day)
	total, _ := den.CalcPayDay(time.Now(), payDay)

	t.Log(total)
}

func TestDenominator_CalcPayDay_Now(t *testing.T) {

	cfg := config.Build()
	den := NewDenominator(cfg.Bank, nil)

	payDay := PayDay{
		Day:    time.Now(),
		Amount: decimal.NewFromInt(100),
	}
	t.Log(payDay.Day)
	total, _ := den.CalcPayDay(time.Now(), payDay)

	t.Log(total)
}
