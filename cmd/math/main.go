package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/files"
	router "github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/slices"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const envAuthKeysName = "AUTH_KEYS_PATHNAME"
const serverSuccessfullyStartedMessage = "Web server is running on port 8080 \n"
const serverDidNotStartMessage = "Web server did not start. Please check the command, should contain --web-server \n"
const command = "go run cmd/uuid/main.go 1000"
const colorNone = "\033[0m"
const colorGreen = "\033[0;32m"

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var (
		authKeysPathname = os.Getenv(envAuthKeysName)
		serverPort       = os.Getenv("SERVER_PORT")
	)

	err := files.FindFile(authKeysPathname)
	if err != nil {
		fmt.Fprintf(os.Stdout, "Keys was not generated yet, please run the command to generate auth keys %s %s %s and try it again", colorGreen, command, colorNone)
		return
	}

	r := router.NewRouter()

	if slices.Contains(os.Args[1:], "--web-server") {
		fmt.Print(serverSuccessfullyStartedMessage)
		log.Fatal(http.ListenAndServe(":"+serverPort, r))
	} else {
		fmt.Print(serverDidNotStartMessage)
	}
}
