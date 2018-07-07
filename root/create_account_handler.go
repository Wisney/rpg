package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
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
	fmt.Println("teste")
	page := strings.Replace(string(formSignup), "{{.Message}}", "Cadastro!", 1)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

// getCreateAccountHandler renders the homepage view template
func postCreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	nick := r.FormValue("nick")
	login := r.FormValue("login")
	password := r.FormValue("password")

	fmt.Println(nick)
	fmt.Println(login)
	fmt.Println(password)

	formSignin, err := ioutil.ReadFile("./pages/form_signin.html")
	if err != nil {
		fmt.Print(err)
	}

	page := strings.Replace(string(formSignin), "{{.Message}}", "Cadastrado com Sucesso!", 1)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}

	render(w, r, homepageTpl, "homepage_view", fullData)
}
