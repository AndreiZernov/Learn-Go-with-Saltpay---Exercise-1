package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/data_retriever"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"os"
)

func main() {
	toGetAllArgs := os.Args[1:]

	dataRetriever := data_retriever.New()
	numbers := dataRetriever.GetData(toGetAllArgs)

	calculator := calculator.New()
	result, err := calculator.Add(numbers)

	formatter := formatter.New()
	formattedResult := formatter.GroupsOfThousands(result)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sum of %s equal %s \n", numbers, formattedResult)
	}
}
