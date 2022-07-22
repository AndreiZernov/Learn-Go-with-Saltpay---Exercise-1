package main

import (
	"fmt"
	http "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp"
	"log"
	"os"
)

func main() {
	toGetAllArgs := os.Args[1:]
	var (
		serverPort  = os.Getenv("SERVER_PORT")
		apiEndpoint = os.Getenv("API_ENDPOINT")
		requestURL  = fmt.Sprintf("%s:%s", apiEndpoint, serverPort)
	)

	fiboClient := http.NewFiboClient(requestURL, nil)

	for _, arg := range toGetAllArgs {
		_, err := fiboClient.Call(arg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
