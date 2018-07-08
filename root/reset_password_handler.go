package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"rpg/infra/db"
	jwt "rpg/infra/security"
)

func getResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	forgotpassword, err := ioutil.ReadFile("./pages/reset_password.html")
	if err != nil {
		fmt.Print(err)
	}

	pathVariables := mux.Vars(r)
	token := pathVariables["token"]

	page := strings.Replace(string(forgotpassword), "{{.Message}}", "Sou noob, esqueci a Senha!", 1)
	page = strings.Replace(string(page), "{{.Warning}}", "", 1)
	page = strings.Replace(string(page), "{{.Token}}", token, 1)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

func postResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	password := r.FormValue("password")

	pathVariables := mux.Vars(r)
	token := pathVariables["token"]

	email, er := jwt.GetEmailFromForgotPasswordToken(token)
	errorPage := false
	if er != nil {
		errorPage = true
	}

	forgotpassword, err := ioutil.ReadFile("./pages/reset_password.html")
	if err != nil {
		fmt.Print(err)
	}

	errorPage = !(errorPage == false && db.UpdatePasswordByEmail(email, password))
	page := string(forgotpassword)
	if !errorPage {
		warning, err := ioutil.ReadFile("./pages/warnings/reset_password_sucess.html")
		if err != nil {
			fmt.Print(err)
		}

		page = strings.Replace(page, "{{.Message}}", "Sou noob, esqueci a Senha!", 1)
		page = strings.Replace(page, "{{.Warning}}", string(warning), 1)

	} else {
		warning, err := ioutil.ReadFile("./pages/warnings/reset_password_invalid.html")
		if err != nil {
			fmt.Print(err)
		}

		page = strings.Replace(page, "{{.Message}}", "Sou noob, esqueci a Senha!", 1)
		page = strings.Replace(page, "{{.Warning}}", string(warning), 1)

	}

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}

	render(w, r, homepageTpl, "homepage_view", fullData)
}
