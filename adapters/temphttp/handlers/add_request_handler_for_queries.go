package handlers

import "net/http"

func AddRequestHandlerForQueries(w http.ResponseWriter, req *http.Request) {
	data := req.URL.Query()["num"]
	AddResponseHandler(w, data)
}
