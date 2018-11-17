package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	jwt "rpg/infra/security"
	"strings"
)

func navbar(w http.ResponseWriter, r *http.Request) {

	//get navbar file
	navigationBarHTMLFile, err := ioutil.ReadFile("./templates/navigation_bar.html")
	if err != nil {
		fmt.Print(err)
	}

	navigationBarHTML = string(navigationBarHTMLFile)

	cookie, err := r.Cookie("Authorization")
	/* if err != nil {
		panic(err)
	} */

	if cookie != nil {
		claims, err := jwt.GetClaimsFromToken(cookie.Value)

		if err != nil || claims["nick"] == nil || claims["nick"].(string) == "" {
			fmt.Println(err)
			goto not_logged
		}

		navBar, err := ioutil.ReadFile("./templates/logged_menu.html")
		if err != nil {
			fmt.Print(err)
		}
		collapsedNavBar, err := ioutil.ReadFile("./templates/collapsed_logged_menu.html")
		if err != nil {
			fmt.Print(err)
		}

		navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.navbar-menu}}", string(navBar), 1)
		navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.collapsed-navbar-menu}}", string(collapsedNavBar), 1)
		navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.name}}", claims["nick"].(string), 2)

		return
	}

not_logged:

	navBar, err := ioutil.ReadFile("./templates/not_logged_menu.html")
	if err != nil {
		fmt.Print(err)
	}
	collapsedNavBar, err := ioutil.ReadFile("./templates/collapsed_not_logged_menu.html")
	if err != nil {
		fmt.Print(err)
	}

	navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.navbar-menu}}", string(navBar), 1)
	navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.collapsed-navbar-menu}}", string(collapsedNavBar), 1)

}

func forceNavbar(w http.ResponseWriter, r *http.Request, nick string) {
	//get navbar file
	navigationBarHTMLFile, err := ioutil.ReadFile("./templates/navigation_bar.html")
	if err != nil {
		fmt.Print(err)
	}

	navigationBarHTML = string(navigationBarHTMLFile)

	if strings.TrimSpace(nick) != "" {

		navBar, err := ioutil.ReadFile("./templates/logged_menu.html")
		if err != nil {
			fmt.Print(err)
		}
		collapsedNavBar, err := ioutil.ReadFile("./templates/collapsed_logged_menu.html")
		if err != nil {
			fmt.Print(err)
		}

		navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.navbar-menu}}", string(navBar), 1)
		navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.collapsed-navbar-menu}}", string(collapsedNavBar), 1)
		navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.name}}", nick, 2)
	} else {

		navBar, err := ioutil.ReadFile("./templates/not_logged_menu.html")
		if err != nil {
			fmt.Print(err)
		}
		collapsedNavBar, err := ioutil.ReadFile("./templates/collapsed_not_logged_menu.html")
		if err != nil {
			fmt.Print(err)
		}

		navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.navbar-menu}}", string(navBar), 1)
		navigationBarHTML = strings.Replace(string(navigationBarHTML), "{{.collapsed-navbar-menu}}", string(collapsedNavBar), 1)
	}
}
