package http

import (
	"encoding/json"
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

type server struct{}

func newRequestHandlers() *server {
	return &server{}
}

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	httpRequestsHandler := newRequestHandlers()

	router.Use(flakinessMiddleware)
	protectedRoutes := router.PathPrefix("/").Subrouter()
	protectedRoutes.Use(VerifyTokenMiddleware, loggingMiddleware)

	protectedRoutes.HandleFunc("/fibonacci/{position}", httpRequestsHandler.fibonacciRequestHandler).Methods(http.MethodGet)
	protectedRoutes.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForQueries).Methods(http.MethodPost).Queries("num", "{[0-9]*?}")
	protectedRoutes.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForFormUrlEncoded).Methods(http.MethodPost).Headers("Content-Type", "application/x-www-form-urlencoded")
	protectedRoutes.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForJson).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	return router
}

func (svr server) fibonacciRequestHandler(w http.ResponseWriter, req *http.Request) {
	trimmedPath := strings.TrimPrefix(req.URL.Path, "/fibonacci/")
	n, err := strconv.ParseInt(trimmedPath, 10, 64)
	error_handler.HandlePanic(err)

	fib := fibonacci.New()
	fibNumber, err := fib.GetNumberFromNumericPosition(n)
	error_handler.HandlePanic(err)

	_, err = fmt.Fprintf(w, "%d \n", fibNumber)
	error_handler.HandlePanic(err)
}

func (svr server) addRequestHandlerForQueries(w http.ResponseWriter, req *http.Request) {
	data := req.URL.Query()["num"]
	svr.addResponseHandler(w, data)
}

func (svr server) addRequestHandlerForFormUrlEncoded(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	error_handler.HandlePanic(err)

	data := req.PostForm["num"]
	svr.addResponseHandler(w, data)
}

func (svr server) addRequestHandlerForJson(w http.ResponseWriter, req *http.Request) {
	var t struct {
		Nums []int
	}
	body, err := io.ReadAll(req.Body)
	err = json.Unmarshal(body, &t)
	error_handler.HandlePanic(err)

	var data []string
	for _, num := range t.Nums {
		data = append(data, strconv.Itoa(num))
	}
	svr.addResponseHandler(w, data)
}

func (svr server) addResponseHandler(w http.ResponseWriter, data []string) {
	numbers := strings.Join(data[:], ",")
	cleanData := strings_helper.DataCleaner(numbers)

	calculate := calculator.New()
	format := formatter.New()

	result, err := calculate.Add(cleanData)
	formattedResult := format.GroupsOfThousands(result)

	_, err = fmt.Fprintf(w, "Sum of %s equal %s \n", numbers, formattedResult)
	error_handler.HandlePanic(err)
}
