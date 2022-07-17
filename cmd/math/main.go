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

func main() {
	http.HandleFunc("/add", New().processes)

	if array_contains.ArrayContains(os.Args[1:], "--web-server") {
		fmt.Print("Web server is running on port 8080 \n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}

type handlers struct{}

func New() *handlers {
	return &handlers{}
}

func (h handlers) processes(w http.ResponseWriter, req *http.Request) {
	headerContentTtype := req.Header.Get("Content-Type")
	q := req.URL.Query()

	if len(q["num"]) > 0 {
		h.response(w, q["num"])
	} else if headerContentTtype == "application/x-www-form-urlencoded" {
		req.ParseForm()
		h.response(w, req.PostForm["num"])
	}
}

func (h handlers) response(w http.ResponseWriter, data []string) {
	numbers := strings.Join(data[:], ",")
	calculator := calculator.New()
	formatter := formatter.New()

	result, err := calculator.Add(numbers)
	formattedResult := formatter.GroupsOfThousands(result)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Fprintf(w, "Sum of %s equal %s \n", numbers, formattedResult)
	}
}
