package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/sum"
)

func main() {
	numbers := []string{"1", "9223372036854775807", "3", "4", "5"}
	result, err := sum.Add(numbers)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sum of %s equal %s \n", numbers, result)
	}
}
