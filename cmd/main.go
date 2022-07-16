package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/commandLine"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	formatter "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers"
)

func main() {
	commandLine := commandLine.New()
	numbers := commandLine.GetArguments()

	result, err := calculator.Add(numbers)
	formattedResult := formatter.Formatter(result)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sum of %s equal %s \n", numbers, formattedResult)
	}
}
