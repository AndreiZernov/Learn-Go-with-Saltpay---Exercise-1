package router

import (
	"encoding/json"
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
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
	router := mux.NewRouter()
	httpRequestsHandler := newRequestHandlers()

	router.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForQueries).Methods(http.MethodPost).Queries("num", "{[0-9]*?}")
	router.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForFormUrlEncoded).Methods(http.MethodPost).Headers("Content-Type", "application/x-www-form-urlencoded")
	router.HandleFunc("/add", httpRequestsHandler.addRequestHandlerForJson).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	return router
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
	calculate := calculator.New()
	format := formatter.New()

	result, err := calculate.Add(numbers)
	formattedResult := format.GroupsOfThousands(result)

	_, err = fmt.Fprintf(w, "Sum of %s equal %s \n", numbers, formattedResult)

	error_handler.HandlePanic(err)
}
