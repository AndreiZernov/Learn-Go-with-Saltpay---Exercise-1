package main

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/array_contains"
	"log"
	"net/http"
	"os"
	"strings"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	numbers := strings.Join(q["num"][:], ",")

	calculator := calculator.New()
	result, err := calculator.Add(numbers)

	formatter := formatter.New()
	formattedResult := formatter.GroupsOfThousands(result)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Fprintf(w, "Sum of %s equal %s \n", numbers, formattedResult)
	}
}

func main() {
	toGetAllArgs := os.Args[1:]

	http.HandleFunc("/add", HelloHandler)

	if array_contains.ArrayContains(toGetAllArgs, "--web-server") {
		fmt.Print("Web server is running on port 8080 \n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}
