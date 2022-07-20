package handlers

import (
	"encoding/json"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"io"
	"net/http"
	"strconv"
)

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
