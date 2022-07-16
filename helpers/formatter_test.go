package formatter_test

import (
	formatter "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers"
	"testing"
)

func TestFormatter(t *testing.T) {
	formaterTest := []struct {
		name            string
		number          string
		formattedNumber string
	}{
		{number: "200000", formattedNumber: "200,000"},
		{number: "-200,000", formattedNumber: "-200,000"},
		{number: "9999", formattedNumber: "9999"},
		{number: "-9999", formattedNumber: "-9999"},
		{number: "999", formattedNumber: "999"},
		{number: "0", formattedNumber: "0"},
	}

	for _, tt := range formaterTest {
		t.Run(tt.name, func(t *testing.T) {
			got := formatter.Formatter(tt.number)
			if tt.formattedNumber != got {
				t.Errorf("expected %s formatted to %s", tt.formattedNumber, got)
			}
		})
	}
}
