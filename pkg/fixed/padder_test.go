package fixed

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestPad(t *testing.T) {
	dec := decimal.RequireFromString("1234567.12345")
	t.Log(Pad(dec))
}
