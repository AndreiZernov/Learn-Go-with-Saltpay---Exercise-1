package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/data_retriever"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"log"
	"os"
)

func main() {
	toGetAllArgs := os.Args[1:]

	dataRetriever := data_retriever.New()
	numbers, err := dataRetriever.GetNumbers(toGetAllArgs)
	if err != nil {
		log.Fatal(err)
	}

	calculate := calculator.New()
	result, err := calculate.Add(numbers)
	if err != nil {
		log.Fatal(err)
	}

	format := formatter.New()
	formattedResult := format.GroupsOfThousands(result, true)

	fmt.Printf("%s\n", formattedResult)
}
