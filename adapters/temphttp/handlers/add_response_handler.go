package handlers

import (
	"errors"
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/helpers/strings_helper"
	"net/http"
	"strings"
)

func AddResponseHandler(w http.ResponseWriter, data []string) {
	if len(data) == 0 {
		err := errors.New("400 Bad Request")
		error_handler.HandleStatusBadRequest(w, err)
		return
	}

	numbers := strings.Join(data[:], ",")
	cleanData := strings_helper.DataCleaner(numbers)

	calculate := calculator.New()
	format := formatter.New()

	result, err := calculate.Add(cleanData)
	formattedResult := format.GroupsOfThousands(result)

	_, err = fmt.Fprintf(w, "Sum of %s equal %s \n", numbers, formattedResult)
	error_handler.HandleStatusBadRequest(w, err)
}
