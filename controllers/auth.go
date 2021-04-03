package controllers

import (
	"net/http"
	"web/middlewares"
	"web/views"

	"github.com/gorilla/mux"
)

func AuthController(r *mux.Router) {
	auth := r.PathPrefix("/auth").Subrouter()
	auth.HandleFunc("/register", middlewares.Chain(views.Register, middlewares.Method(http.MethodGet), middlewares.Logging()))
	auth.HandleFunc("/login", middlewares.Chain(views.Login, middlewares.Method(http.MethodGet), middlewares.Logging()))
}
