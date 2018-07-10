package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	db "rpg/infra/db"
	"strings"
)

// getCreateAccountHandler renders the homepage view template
func getCreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	formSignup, err := ioutil.ReadFile("./pages/form_signup.html")
	if err != nil {
		fmt.Print(err)
	}

	page := strings.Replace(string(formSignup), "{{.Message}}", "Cadastro!", 1)
	page = strings.Replace(string(page), "{{.NickError}}", "", 1)
	page = strings.Replace(string(page), "{{.EmailError}}", "", 1)
	page = strings.Replace(string(page), "{{.PasswordError}}", "", 1)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

// postCreateAccountHandler renders the homepage view template
func postCreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	nick := r.FormValue("nick")
	email := r.FormValue("email")
	password := r.FormValue("password")

	fmt.Println(nick)
	fmt.Println(email)
	fmt.Println(password)

	nickError := ""
	emailError := ""
	passError := ""

	errorPage := false
	if db.ExistNick(nick) {
		nickErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
		if err != nil {
			fmt.Print(err)
		}
		errorString := string(nickErrorByte) + string(nickErrorByte)

		nickError = strings.Replace(errorString, "{{.Message}}", nick, 1)
		nickError = strings.Replace(nickError, "{{.Message}}", "Já está sendo usado!", 1)
		errorPage = true
	}

	if db.ExistEmail(email) {
		emailErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
		if err != nil {
			fmt.Print(err)
		}
		errorString := string(emailErrorByte) + string(emailErrorByte)

		emailError = strings.Replace(errorString, "{{.Message}}", email, 1)
		emailError = strings.Replace(emailError, "{{.Message}}", "Já está sendo usado!", 1)
		errorPage = true
	}

	if strings.TrimSpace(password) == "" {
		passwordErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
		if err != nil {
			fmt.Print(err)
		}
		errorString := string(passwordErrorByte)

		passError += strings.Replace(errorString, "{{.Message}}", "Password vazio! ¬¬", 1)
		errorPage = true
	}

	if !errorPage {
		home, err := ioutil.ReadFile("./pages/home.html")
		if err != nil {
			fmt.Print(err)
		}

		page := strings.Replace(string(home), "{{.Message}}", "Logado!", 1)

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(navigationBarHTML),
			"Page":          template.HTML(page),
		}

		render(w, r, homepageTpl, "homepage_view", fullData)
	} else {

		formSignup, err := ioutil.ReadFile("./pages/form_signup.html")
		if err != nil {
			fmt.Print(err)
		}

		page := strings.Replace(string(formSignup), "{{.Message}}", "Cadastro!", 1)
		page = strings.Replace(string(page), "{{.NickError}}", nickError, 1)
		page = strings.Replace(string(page), "{{.EmailError}}", emailError, 1)
		page = strings.Replace(string(page), "{{.PasswordError}}", passError, 1)

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(navigationBarHTML),
			"Page":          template.HTML(page),
		}
		render(w, r, homepageTpl, "homepage_view", fullData)
	}
}
