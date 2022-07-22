package main

import (
	http "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp"
	"os"
)

func main() {
	toGetAllArgs := os.Args[1:]

	fiboClient := http.NewFiboClient()

	for _, arg := range toGetAllArgs {
		fiboClient.Call(arg)
	}
}
