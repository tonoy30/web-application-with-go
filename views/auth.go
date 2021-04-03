package views

import (
	"net/http"
	"web/utils"
	"web/viewmodels"
)

func Login(w http.ResponseWriter, r *http.Request) {
	utils.ServeTemplate(viewmodels.NewHead("Login"), "login", w)
}
func Register(w http.ResponseWriter, r *http.Request) {
	utils.ServeTemplate(viewmodels.NewHead("Register"), "signup", w)
}
