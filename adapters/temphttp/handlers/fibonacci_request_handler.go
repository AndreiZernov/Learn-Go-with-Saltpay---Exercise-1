package handlers

import (
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/fibonacci"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"net/http"
	"strconv"
	"strings"
)

func FibonacciRequestHandler(w http.ResponseWriter, req *http.Request) {
	formatQuery := req.URL.Query()["format"]

	trimmedPath := strings.TrimPrefix(req.URL.Path, "/fibonacci/")
	position, err := strconv.ParseInt(trimmedPath, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fib := fibonacci.New()
	fibNumber, err := fib.GetNumberFromNumericPosition(position)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	format := formatter.New()
	stringFibNumber := format.GroupsOfThousands(fibNumber, len(formatQuery) > 0 && formatQuery[0] == "thousands")

	_, err = fmt.Fprintf(w, "%s", stringFibNumber)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
