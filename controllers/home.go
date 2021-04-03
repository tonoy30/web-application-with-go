package controllers

import (
	"net/http"
	"web/middlewares"
	"web/views"

	"github.com/gorilla/mux"
)

func HomeController(r *mux.Router) {
	r.HandleFunc("/", middlewares.Chain(views.Hello, middlewares.Method(http.MethodGet), middlewares.Logging()))
	r.HandleFunc("/service", middlewares.Chain(views.Service, middlewares.Method(http.MethodGet), middlewares.Logging()))
}
