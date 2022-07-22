package main

import (
	http "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp"
	"log"
	"os"
)

func main() {
	toGetAllArgs := os.Args[1:]

	fiboClient := http.NewFiboClient()

	for _, arg := range toGetAllArgs {
		err := fiboClient.Call(arg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
