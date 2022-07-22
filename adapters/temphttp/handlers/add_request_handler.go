package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/data_retriever"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/calculator"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/domain/formatter"
	"html"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func AddRequestHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	numQueries := req.URL.Query()["num"]
	formatQuery := req.URL.Query()["format"]
	var data []string

	switch {
	case len(numQueries) > 0:
		data = numQueries

	case strings.Contains(contentType, "application/x-www-form-urlencoded"):
		err := req.ParseForm()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data = req.PostForm["num"]

	case strings.Contains(contentType, "application/json"):
		var jsonBody struct {
			Nums []int
		}
		body, err := io.ReadAll(req.Body)
		err = json.Unmarshal(body, &jsonBody)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		for _, num := range jsonBody.Nums {
			data = append(data, strconv.Itoa(num))
		}
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dataRetriever := data_retriever.New()
	numbers, err := dataRetriever.GetNumbers(data)

	calculate := calculator.New()
	format := formatter.New()

	result, err := calculate.Add(numbers)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	formattedResult := format.GroupsOfThousands(result, len(formatQuery) > 0 && formatQuery[0] == "thousands")
	responseMessage := fmt.Sprintf("%s", formattedResult)

	_, err = fmt.Fprintf(w, "%s", html.EscapeString(responseMessage))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
