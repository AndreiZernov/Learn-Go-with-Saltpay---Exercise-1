package main

import (
	"fmt"
	router "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/http"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	r := router.NewRouter()

	if slices.Contains(os.Args[1:], "--web-server") {
		fmt.Print("Web server is running on port 8080 \n")
		log.Fatal(http.ListenAndServe(":8080", r))
	} else {
		fmt.Print("Web server did not start. Please check the command, should contain --web-server \n")
	}
}
