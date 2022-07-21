package temphttp

import (
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/handlers"
	"github.com/AndreiZernov/learn_go_with_saltpay_exercise_one/adapters/temphttp/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Use(middlewares.FlakinessMiddleware)
	protectedRoutes := router.PathPrefix("/").Subrouter()
	protectedRoutes.Use(middlewares.AuthenticationMiddleware, middlewares.LoggingMiddleware)

	protectedRoutes.HandleFunc("/fibonacci/{position}", handlers.FibonacciRequestHandler).Methods(http.MethodGet)
	protectedRoutes.HandleFunc("/add", handlers.AddRequestHandler).Methods(http.MethodPost)

	return router
}
