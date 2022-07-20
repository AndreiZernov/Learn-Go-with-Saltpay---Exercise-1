package handlers

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/fibonacci"
	"net/http"
	"strconv"
	"strings"
)

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
