package main

import (
	"log"
	"net/http"
	"web/controllers"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8000"
)

func main() {
	r := mux.NewRouter()
	r.Headers("x-content-type-options", "nosniff")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	controllers.HomeController(r)
	controllers.AuthController(r)
	log.Fatal(http.ListenAndServe(PORT, r))
}
