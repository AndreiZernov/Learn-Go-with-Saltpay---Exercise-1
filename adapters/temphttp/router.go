package temphttp

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(flakinessMiddleware)
	protectedRoutes := router.PathPrefix("/").Subrouter()
	protectedRoutes.Use(VerifyTokenMiddleware, loggingMiddleware)

	protectedRoutes.HandleFunc("/fibonacci/{position}", handlers.FibonacciRequestHandler).Methods(http.MethodGet)
	protectedRoutes.HandleFunc("/add", handlers.AddRequestHandlerForQueries).Methods(http.MethodPost).Queries("num", "{[0-9]*?}")
	protectedRoutes.HandleFunc("/add", handlers.AddRequestHandlerForFormUrlEncoded).Methods(http.MethodPost).Headers("Content-Type", "application/x-www-form-urlencoded")
	protectedRoutes.HandleFunc("/add", handlers.AddRequestHandlerForJson).Methods(http.MethodPost).Headers("Content-Type", "application/json")

	return router
}
