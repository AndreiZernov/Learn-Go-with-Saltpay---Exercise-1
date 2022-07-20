package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/data_retriever"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/strings_helper"
	"os"
)

func main() {
	toGetAllArgs := os.Args[1:]

	dataRetriever := data_retriever.New()
	numbers := dataRetriever.GetData(toGetAllArgs)

	cleanData := strings_helper.DataCleaner(numbers)

	calculate := calculator.New()
	result, err := calculate.Add(cleanData)

	format := formatter.New()
	formattedResult := format.GroupsOfThousands(result, true)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("Sum of %s equal %s \n", cleanData, formattedResult)
	}
}
