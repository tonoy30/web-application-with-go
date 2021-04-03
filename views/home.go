package views

import (
	"net/http"
	"web/utils"
	"web/viewmodels"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	utils.ServeTemplate(viewmodels.NewHead("Home"), "home", w)
}

func Service(w http.ResponseWriter, r *http.Request) {
	utils.ServeTemplate(viewmodels.NewHead("Service"), "service", w)
}
