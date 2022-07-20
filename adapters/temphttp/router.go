package temphttp

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/fibonacci"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/strings_helper"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(flakinessMiddleware)
	protectedRoutes := router.PathPrefix("/").Subrouter()
	protectedRoutes.Use(VerifyTokenMiddleware, loggingMiddleware)

	protectedRoutes.HandleFunc("/fibonacci/{position}", FibonacciRequestHandler).Methods(http.MethodGet)
	protectedRoutes.HandleFunc("/add", AddRequestHandlerForQueries).Methods(http.MethodPost).Queries("num", "{[0-9]*?}")
	protectedRoutes.HandleFunc("/add", AddRequestHandlerForFormUrlEncoded).Methods(http.MethodPost).Headers("Content-Type", "application/x-www-form-urlencoded")
	protectedRoutes.HandleFunc("/add", AddRequestHandlerForJson).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	return router
}

func FibonacciRequestHandler(w http.ResponseWriter, req *http.Request) {
	trimmedPath := strings.TrimPrefix(req.URL.Path, "/fibonacci/")
	n, err := strconv.ParseInt(trimmedPath, 10, 64)
	error_handler.HandleStatusBadRequest(w, err)

	fib := fibonacci.New()
	fibNumber, err := fib.GetNumberFromNumericPosition(n)
	error_handler.HandleStatusBadRequest(w, err)

	_, err = fmt.Fprintf(w, "%d \n", fibNumber)
	error_handler.HandleStatusBadRequest(w, err)
}

func AddRequestHandlerForQueries(w http.ResponseWriter, req *http.Request) {
	data := req.URL.Query()["num"]
	if len(data) == 0 {
		err := errors.New("400 Bad Request")
		error_handler.HandleStatusBadRequest(w, err)
	} else {
		AddResponseHandler(w, data)
	}
}

func AddRequestHandlerForFormUrlEncoded(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	error_handler.HandleStatusBadRequest(w, err)

	data := req.PostForm["num"]
	AddResponseHandler(w, data)
}

func AddRequestHandlerForJson(w http.ResponseWriter, req *http.Request) {
	var t struct {
		Nums []int
	}
	body, err := io.ReadAll(req.Body)
	err = json.Unmarshal(body, &t)
	error_handler.HandleStatusBadRequest(w, err)

	var data []string
	for _, num := range t.Nums {
		data = append(data, strconv.Itoa(num))
	}
	AddResponseHandler(w, data)
}

func AddResponseHandler(w http.ResponseWriter, data []string) {
	numbers := strings.Join(data[:], ",")
	cleanData := strings_helper.DataCleaner(numbers)

	calculate := calculator.New()
	format := formatter.New()

	result, err := calculate.Add(cleanData)
	formattedResult := format.GroupsOfThousands(result)

	_, err = fmt.Fprintf(w, "Sum of %s equal %s \n", numbers, formattedResult)
	error_handler.HandleStatusBadRequest(w, err)
}
