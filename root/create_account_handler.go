package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	db "rpg/infra/db"
	"strings"
	"time"
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
	page = strings.Replace(string(page), "{{.CaptchaError}}", "", 1)

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
	recaptcha := r.FormValue("g-recaptcha-response")

	/* fmt.Println(nick)
	fmt.Println(email)
	fmt.Println(password)
	fmt.Println(recaptcha) */

	validCaptcha := reCaptchaRequest(recaptcha)

	nickError := ""
	emailError := ""
	passError := ""
	captchaError := ""

	errorPage := false
	if db.ExistNick(nick) {
		nickErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
		if err != nil {
			fmt.Print(err)
		}
		errorString := string(nickErrorByte) + string(nickErrorByte)

		nickError = strings.Replace(errorString, "{{.Error}}", nick, 1)
		nickError = strings.Replace(nickError, "{{.Error}}", "Já está sendo usado!", 1)
		errorPage = true
	}

	if db.ExistEmail(email) {
		emailErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
		if err != nil {
			fmt.Print(err)
		}
		errorString := string(emailErrorByte) + string(emailErrorByte)

		emailError = strings.Replace(errorString, "{{.Error}}", email, 1)
		emailError = strings.Replace(emailError, "{{.Error}}", "Já está sendo usado!", 1)
		errorPage = true
	}

	if strings.TrimSpace(password) == "" {
		passwordErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
		if err != nil {
			fmt.Print(err)
		}
		errorString := string(passwordErrorByte)

		passError += strings.Replace(errorString, "{{.Error}}", "Password vazio! ¬¬", 1)
		errorPage = true
	}

	if !validCaptcha {
		captchaErrorByte, err := ioutil.ReadFile("./pages/warnings/generic_invalid.html")
		if err != nil {
			fmt.Print(err)
		}
		errorString := string(captchaErrorByte)

		captchaError += strings.Replace(errorString, "{{.Error}}", "Marque o reCAPTCHA!", 1)
		errorPage = true
	}

	if !errorPage {
		//signup and create token

		home, err := ioutil.ReadFile("./pages/home.html")
		if err != nil {
			fmt.Print(err)
		}

		page := strings.Replace(string(home), "{{.Message}}", "Logado!", 1)

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(navigationBarHTML),
			"Page":          template.HTML(page),
		}

		//Set Authorization token
		expire := time.Now().Add(7 * 24 * time.Hour) // Expires in 7 days
		cookie := http.Cookie{Name: "Authorization", Value: "test", Path: "/", Expires: expire, MaxAge: 604800, HttpOnly: true, Secure: false}
		http.SetCookie(w, &cookie)

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
		page = strings.Replace(string(page), "{{.CaptchaError}}", captchaError, 1)

		fullData := map[string]interface{}{
			"NavigationBar": template.HTML(navigationBarHTML),
			"Page":          template.HTML(page),
		}

		render(w, r, homepageTpl, "homepage_view", fullData)
	}
}

//reCaptchaRequest send a post to google/recaptcha to verify integrity
func reCaptchaRequest(captcha string) bool {

	formData := url.Values{
		"secret":   {os.Getenv("RECAPTCHA")},
		"response": {captcha},
	}

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", formData)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)

	if result["success"] == nil {
		return false
	}

	return result["success"].(bool)
}
