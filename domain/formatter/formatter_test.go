package formatter_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatter(t *testing.T) {
	formaterTest := []struct {
		name            string
		number          int
		formattedNumber string
	}{
		{name: "Given a number bigger than 9999 should return formatted number with comma for groups of thousands", number: 200000, formattedNumber: "200,000"},
		{name: "Given a number smaller than -9999 should return formatted number with comma for groups of thousands", number: -200000, formattedNumber: "-200,000"},
		{name: "Given a number equal 9999 should return the same number", number: 9999, formattedNumber: "9999"},
		{name: "Given a number equal -9999 should return the same number", number: -9999, formattedNumber: "-9999"},
		{name: "Given a number smaller than 9999 and bigger than -9999 should return the same number", number: 10, formattedNumber: "10"},
		{name: "Given a number equal 0 should return 0", number: 0, formattedNumber: "0"},
	}

	for _, tt := range formaterTest {
		t.Run(tt.name, func(t *testing.T) {
			formatter := formatter.New()
			got := formatter.GroupsOfThousands(tt.number)
			assert.Equal(t, tt.formattedNumber, got)
		})
	}
}
