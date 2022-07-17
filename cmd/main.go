package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/commandLine"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
)

func main() {
	numbers := commandLine.GetArguments()
	calculator := calculator.New()

	var result, summary int
	var err error

	for _, number := range numbers {
		summary, err = calculator.Add(number)
		result += summary
	}

	formatter := formatter.New()
	formattedResult := formatter.GroupsOfThousands(result)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sum of %s equal %s \n", numbers, formattedResult)
	}
}
