package main

import (
	"encoding/json"
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/array_contains"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/add", new().processes)

	if array_contains.ArrayContains(os.Args[1:], "--web-server") {
		fmt.Print("Web server is running on port 8080 \n")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		fmt.Print("Web server did not start. Please check the command, should contain --web-server \n")
	}
}

type handlers struct{}

func new() *handlers {
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
	} else if headerContentTtype == "application/json" {
		data := h.RetrieveJsonArray(req)
		h.response(w, data)
	} else {
		w.WriteHeader(http.StatusBadRequest)
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

func (h handlers) RetrieveJsonArray(req *http.Request) []string {
	var t struct {
		Nums []int
	}
	body, err := io.ReadAll(req.Body)
	err = json.Unmarshal(body, &t)
	if err != nil {
		panic(err)
	}

	var data []string
	for _, num := range t.Nums {
		data = append(data, strconv.Itoa(num))
	}
	return data
}
