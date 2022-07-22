package formatter_test

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatter(t *testing.T) {
	formaterTest := []struct {
		Name            string
		Number          int64
		FormattedNumber string
	}{
		{
			Name:            "Given a Number bigger than 9999 should return formatted Number with comma for groups of thousands",
			Number:          200000,
			FormattedNumber: "200,000",
		},
		{
			Name:            "Given a Number smaller than -9999 should return formatted Number with comma for groups of thousands",
			Number:          -200000,
			FormattedNumber: "-200,000",
		},
		{
			Name:            "Given a Number equal 9999 should return the same Number",
			Number:          9999,
			FormattedNumber: "9999",
		},
		{
			Name:            "Given a Number equal -9999 should return the same Number",
			Number:          -9999,
			FormattedNumber: "-9999",
		},
		{
			Name:            "Given a Number smaller than 9999 and bigger than -9999 should return the same Number",
			Number:          10,
			FormattedNumber: "10",
		},
		{
			Name:            "Given a Number equal 0 should return 0",
			Number:          0,
			FormattedNumber: "0",
		},
	}

	for _, tt := range formaterTest {
		t.Run(tt.Name, func(t *testing.T) {
			format := formatter.New()
			got := format.GroupsOfThousands(tt.Number, true)
			assert.Equal(t, tt.FormattedNumber, got)
		})
	}

	t.Run("Given the format boolean false should return the same number", func(t *testing.T) {
		format := formatter.New()
		got := format.GroupsOfThousands(10000000, false)
		assert.Equal(t, "10000000", got)
	})
}
