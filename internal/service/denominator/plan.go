package denominator

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

type Plan struct {
	Start   time.Time
	PayDays []PayDay
}

type PayDay struct {
	Day    time.Time
	Amount decimal.Decimal
}

func NewPlanFromRaw(msg string) (Plan, error) {

	lines := strings.Split(msg, "\n")
	payDays := make([]PayDay, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)

		parts := strings.SplitN(line, " ", 2)
		if len(parts) != 2 {
			parts = strings.SplitN(line, "\t", 2)
			if len(parts) != 2 {
				return Plan{}, fmt.Errorf("invalid plan line: %s", line)
			}
		}

		day, err := parseDay(parts[0])
		if err != nil {
			return Plan{}, fmt.Errorf("invalid plan date: %s", line)
		}

		dec, err := parseDecimal(parts[1])
		if err != nil {
			return Plan{}, fmt.Errorf("invalid plan decimal: %s", line)
		}

		payDays = append(payDays, PayDay{
			Day:    day,
			Amount: dec,
		})
	}

	plan := Plan{
		Start:   time.Now(),
		PayDays: payDays,
	}

	return plan, nil
}

func parseDay(text string) (time.Time, error) {
	text = strings.Trim(text, " ,;\t")

	layouts := []string{
		time.DateOnly,
		"2006.01.02",
		"2006,01,02",
		"2006/01/02",
		"02-01-2006",
		"02.01.2006",
		"02,01,2006",
		"02/01/2006",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, text)
		if err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid day: %s", text)
}

func parseDecimal(text string) (decimal.Decimal, error) {
	text = strings.Trim(text, " ,;\t")
	text = strings.ReplaceAll(text, " ", "")
	text = strings.ReplaceAll(text, "_", "")
	text = strings.ReplaceAll(text, ",", ".")

	dec, err := decimal.NewFromString(text)
	if err != nil {
		return decimal.Decimal{}, fmt.Errorf("invalid decimal: %s", text)
	}

	return dec, nil
}

func (p Plan) Valid() error {
	for _, day := range p.PayDays {
		if p.Start.After(day.Day) {
			return fmt.Errorf("start day after: %s", day.Day)
		}
	}

	if len(p.PayDays) == 0 {
		return fmt.Errorf("no valid days")
	}

	prev := p.PayDays[0]
	for i := 1; i < len(p.PayDays); i++ {
		if p.PayDays[i].Day.Before(prev.Day) {
			return fmt.Errorf("wrong sequense: %s -> %s", prev.Day, p.PayDays[i].Day)
		}
	}

	return nil
}
