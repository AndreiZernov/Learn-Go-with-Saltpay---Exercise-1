package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/sum"
	formatter "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers"
)

func main() {
	numbers := []string{"1", "1", "3", "4", "5"}
	result, err := sum.Add(numbers)
	formatter.Formatter(result)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sum of %s equal %s \n", numbers, result)
	}
}
