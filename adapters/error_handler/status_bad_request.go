package error_handler

import "net/http"

func HandleStatusBadRequest(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
