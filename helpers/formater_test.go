package formater_test

import (
	formater "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers"
	"testing"
)

func TestFormater(t *testing.T) {
	t.Run("should return a string with the correct format", func(t *testing.T) {
		got := formater.Formater("200000")
		want := "200,000"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
