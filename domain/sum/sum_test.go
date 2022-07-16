package sum_test

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/sum"
	"testing"
)

func TestAdder(t *testing.T) {
	adderTest := []struct {
		name    string
		numbers []string
		sum     string
	}{
		{numbers: []string{"0", "0"}, sum: "0"},
		{numbers: []string{"1", "2", "3"}, sum: "6"},
		{numbers: []string{"-7", "8", "-9"}, sum: "-8"},
		{numbers: []string{"add", "2", "2"}, sum: "4"},
		{numbers: []string{"0", "2.4", "2"}, sum: "2"},
		{numbers: []string{"3", "9223372036854775807"}, sum: "0"},
		{numbers: []string{"-2", "-9223372036854775808"}, sum: "0"},
	}

	for _, tt := range adderTest {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sum.Add(tt.numbers)
			checkAdder(t, err, tt.sum, got)
		})
	}
}

func checkAdder(t testing.TB, err error, expected, sum string) {
	t.Helper()
	if err != nil {
		fmt.Println(err)
	}
	if expected != sum {
		t.Errorf("expected %s sum %s", expected, sum)
	}
}
