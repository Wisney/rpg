package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"rpg/infra/db"
	"strings"
)
import jwt "rpg/infra/security"

func getHomeHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	navbar(w, r)

	if jwt.IsValidCookie(r) {
		char, err := ioutil.ReadFile("./pages/character_sheet.html")
		if err != nil {
			fmt.Print(err)
		}

		page := string(char)
		//page := strings.Replace(string(char), "{{.Message}}", "Bem Vindo!", 1)

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(navigationBarHTML),
			"Page":          template.HTML(page),
		}

		render(w, r, homepageTpl, "homepage_view", fullData)
	} else {

		formSignin, err := ioutil.ReadFile("./pages/form_signin.html")
		if err != nil {
			fmt.Print(err)
		}

		page := strings.Replace(string(formSignin), "{{.LoginError}}", "", 1)

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(navigationBarHTML),
			"Page":          template.HTML(page),
		}

		render(w, r, homepageTpl, "homepage_view", fullData)
	}
}

func postHomeHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	navbar(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	nick := r.FormValue("nick")
	password := r.FormValue("password")

	loginError := ""

	errorPage := false
	if strings.TrimSpace(password) == "" || strings.TrimSpace(nick) == "" {
		passwordErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
		if err != nil {
			fmt.Print(err)
		}
		errorString := string(passwordErrorByte)

		loginError = strings.Replace(errorString, "{{.Error}}", "Nick ou Senha Vazios! ¬¬", 1)
		errorPage = true
	} else {

		token, err := db.Signin(nick, password)

		if err != nil {
			nickErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
			if err != nil {
				fmt.Print(err)
			}
			errorString := string(nickErrorByte)

			loginError = strings.Replace(errorString, "{{.Error}}", "Nick ou Senha Inválido!", 1)
			errorPage = true
		} else {
			//Set Authorization token
			jwt.SetCookieToken(w, token)

			nick, err = jwt.GetNickFromToken(token)
		}
	}

	if !errorPage {
		home, err := ioutil.ReadFile("./pages/home.html")
		if err != nil {
			fmt.Print(err)
		}

		page := strings.Replace(string(home), "{{.Message}}", "Bem Vindo!", 1)

		forceNavbar(w, r, nick)

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(navigationBarHTML),
			"Page":          template.HTML(page),
		}

		render(w, r, homepageTpl, "homepage_view", fullData)
	} else {

		formSignup, err := ioutil.ReadFile("./pages/form_signin.html")
		if err != nil {
			fmt.Print(err)
		}

		page := strings.Replace(string(formSignup), "{{.LoginError}}", loginError, 1)
		page = strings.Replace(string(page), "{{.Message}}", "Welcome!", 1)

		forceNavbar(w, r, nick)
		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(navigationBarHTML),
			"Page":          template.HTML(page),
		}

		render(w, r, homepageTpl, "homepage_view", fullData)
	}
}

func getLogoutHomeHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	navbar(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	jwt.SetExpiredCookie(w)

	//get html file
	formSignin, err := ioutil.ReadFile("./pages/form_signin.html")
	if err != nil {
		fmt.Print(err)
	}

	page := strings.Replace(string(formSignin), "{{.LoginError}}", "", 1)

	forceNavbar(w, r, "")
	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}

	render(w, r, homepageTpl, "homepage_view", fullData)
}
