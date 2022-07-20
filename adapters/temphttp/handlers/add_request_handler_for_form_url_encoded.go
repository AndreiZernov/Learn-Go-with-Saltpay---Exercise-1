package handlers

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/error_handler"
	"net/http"
)

func AddRequestHandlerForFormUrlEncoded(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	error_handler.HandleStatusBadRequest(w, err)

	data := req.PostForm["num"]
	AddResponseHandler(w, data)
}
