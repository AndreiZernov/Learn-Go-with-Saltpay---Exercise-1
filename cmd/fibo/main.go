package main

import (
	http "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/http"
)

func main() {
	fiboClient := http.NewFiboClient()
	fiboClient.Call()
}
