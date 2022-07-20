package error_handler

import "net/http"

func HandlePanic(err error) {
	if err != nil {
		panic(err)
	}
}

func HandleStatusBadRequest(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
