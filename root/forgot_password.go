package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
)
import db "rpg/infra/db"
import emailManager "rpg/infra/email"
import jwt "rpg/infra/security"

// getForgotPasswordHandler renders the homepage view template
func getForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	forgotpassword, err := ioutil.ReadFile("./pages/forgot_password.html")
	if err != nil {
		fmt.Print(err)
	}

	page := strings.Replace(string(forgotpassword), "{{.Message}}", "Sou noob, esqueci a Senha!", 1)
	page = strings.Replace(string(page), "{{.Warning}}", "", 1)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

// postForgotPasswordHandler renders the homepage view template
func postForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	push(w, "/static/navigation_bar.css")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	email := r.FormValue("email")

	fmt.Println(email)

	forgotpassword, err := ioutil.ReadFile("./pages/forgot_password.html")
	if err != nil {
		fmt.Print(err)
	}

	page := string(forgotpassword)
	if db.ExistEmail(email) {
		token := jwt.GenerateForgotPasswordToken(email)
		emailManager.SendForgotPasswordEmail(email, token)
		warning, err := ioutil.ReadFile("./pages/warnings/email_sended.html")
		if err != nil {
			fmt.Print(err)
		}

		page = strings.Replace(page, "{{.Message}}", "Sou noob, esqueci a Senha!", 1)
		page = strings.Replace(page, "{{.Warning}}", string(warning), 1)
	} else {
		warning, err := ioutil.ReadFile("./pages/warnings/email_invalid.html")
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
