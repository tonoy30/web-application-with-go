package views

import (
	"net/http"
	"web/utils"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	utils.ServeTemplate(nil, "home", w)
}

func Service(w http.ResponseWriter, r *http.Request) {
	utils.ServeTemplate(nil, "service", w)
}
