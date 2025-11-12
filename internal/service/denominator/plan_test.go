package denominator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseDay(t *testing.T) {

	valids := []string{
		"2025.10.26",
		"2025,10,26",
		"2025/10/26",
		"2025-10-26",
		"26,10,2025",
		"26-10-2025",
		"26/10/2025",
	}

	for _, valid := range valids {
		parsed, err := parseDay(valid)
		assert.NoError(t, err)
		assert.Equal(t, 2025, parsed.Year())
		assert.Equal(t, 10, int(parsed.Month()))
		assert.Equal(t, 26, parsed.Day())

	}
}

func TestName(t *testing.T) {
	txt := "date 1 00 10 010. 43"
	//txt := "datesdfg"
	parts := strings.SplitN(txt, " ", 2)
	t.Log(len(parts))
	t.Log(fmt.Sprintf("'%s'", parts[0]))
	t.Log(fmt.Sprintf("'%s'", parts[1]))
}
