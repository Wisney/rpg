package main

import (
	"crypto/sha256"
	"encoding/base64"
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
	password := encryptPassword(r.FormValue("password"))

	fmt.Println(nick)
	fmt.Println(email)
	fmt.Println(password)

	//maybe login automatic here is better...
	formSignin, err := ioutil.ReadFile("./pages/form_signin.html")
	if err != nil {
		fmt.Print(err)
	}

	page := strings.Replace(string(formSignin), "{{.Message}}", "Cadastrado com Sucesso!<br> Logue agora!", 1)

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(page),
	}

	render(w, r, homepageTpl, "homepage_view", fullData)
}

func encryptPassword(password string) string {
	h := sha256.Sum256([]byte(password))
	return base64.StdEncoding.EncodeToString(h[:])
}
