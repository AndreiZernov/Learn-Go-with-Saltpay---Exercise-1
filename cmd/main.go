package main

import (
	"fmt"
	sum "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one"
)

func main() {
	numbers := []string{"1", "9223372036854775807", "3", "4", "5"}
	result, err := sum.Add(numbers)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sum of %s equal %d \n", numbers, result)
	}
}
