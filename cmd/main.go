package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/sum"
	formatter "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers"
)

func main() {
	numbers := "add, 12121212 , 3, 4, 5"
	result, err := sum.Add(numbers)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		formattedResult := formatter.Formatter(result)
		fmt.Printf("Sum of %s equal %s \n", numbers, formattedResult)
	}
}
