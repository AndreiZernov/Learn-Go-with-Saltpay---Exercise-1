package handlers

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/fibonacci"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"net/http"
	"strconv"
	"strings"
)

func FibonacciRequestHandler(w http.ResponseWriter, req *http.Request) {
	formatQuery := req.URL.Query()["format"]

	trimmedPath := strings.TrimPrefix(req.URL.Path, "/fibonacci/")
	n, err := strconv.ParseInt(trimmedPath, 10, 64)
	error_handler.HandleStatusBadRequest(w, err)

	fib := fibonacci.New()
	fibNumber, err := fib.GetNumberFromNumericPosition(n)
	error_handler.HandleStatusBadRequest(w, err)

	format := formatter.New()
	stringFibNumber := format.GroupsOfThousands(int(fibNumber), len(formatQuery) > 0 && formatQuery[0] == "thousands")

	_, err = fmt.Fprintf(w, "%s", stringFibNumber)
	error_handler.HandleStatusBadRequest(w, err)
}
