package main

import (
	"encoding/json"
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/array_contains"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func main() {
	router := mux.NewRouter()
	httpRequestsHandler := newRequestHandlers()

	router.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForQueries).Methods(http.MethodPost).Queries("num", "{[0-9]*?}")
	router.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForFormUrlEncoded).Methods(http.MethodPost).Headers("Content-Type", "application/x-www-form-urlencoded")
	router.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForJson).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	if array_contains.ArrayContains(os.Args[1:], "--web-server") {
		fmt.Print("Web server is running on port 8080 \n")
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		fmt.Print("Web server did not start. Please check the command, should contain --web-server \n")
	}
}

type handlers struct{}

func newRequestHandlers() *handlers {
	return &handlers{}
}

func (h handlers) addRequestHandlerForQueries(w http.ResponseWriter, req *http.Request) {
	data := req.URL.Query()["num"]
	h.addResponseHandler(w, data)
}

func (h handlers) addRequestHandlerForFormUrlEncoded(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}
	data := req.PostForm["num"]
	h.addResponseHandler(w, data)
}

func (h handlers) addRequestHandlerForJson(w http.ResponseWriter, req *http.Request) {
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
	h.addResponseHandler(w, data)
}

func (h handlers) addResponseHandler(w http.ResponseWriter, data []string) {
	numbers := strings.Join(data[:], ",")
	calculate := calculator.New()
	format := formatter.New()

	result, err := calculate.Add(numbers)
	formattedResult := format.GroupsOfThousands(result)

	_, err = fmt.Fprintf(w, "Sum of %s equal %s \n", numbers, formattedResult)

	if err != nil {
		panic(err)
	}
}
