package fixed

import (
	"github.com/shopspring/decimal"
	"strings"
)

func Pad(num decimal.Decimal) string {
	num = num.Round(2)
	if num.LessThan(decimal.NewFromInt(1000)) {
		return num.String()
	}

	txt := num.String()
	point := len(txt) - 1
	if idx := strings.Index(txt, "."); idx != -1 {
		point = idx
	}

	padded := strings.Builder{}
	//padded.WriteString(txt[point:])
	//padded.String()
	for i, ch := range txt {

		if (point-i)%3 != 0 {
			padded.WriteRune(ch)
			continue
		}

		if i >= point {
			padded.WriteRune(ch)
			continue
		}

		padded.WriteRune(' ')
		padded.WriteRune(ch)
	}

	return padded.String()
}
