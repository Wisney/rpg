package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	db "rpg/infra/db"
	jwt "rpg/infra/security"
	"strconv"
	"time"
)

// getCharacterHandler renders the Character Sheet view template
func getCharacterHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	navbar(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	characterSheet, err := ioutil.ReadFile("./pages/character_sheet.html")
	if err != nil {
		fmt.Print(err)
	}

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(characterSheet),
	}
	render(w, r, homepageTpl, "homepage_view", fullData)
}

// postCharacterHandler Save the Character Sheet informations
func postCharacterHandler(w http.ResponseWriter, r *http.Request) {
	push(w, "/static/style.css")
	navbar(w, r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := r.ParseForm(); err != nil {
		fmt.Println("Error on parseForm!")
		fmt.Println(err)
	}

	character := db.GetCharacter(jwt.GetNickFromRequest(r))

	character.Name = r.FormValue("Name")

	character.Description = r.FormValue("Description")

	character.ID = getInt8FromString(r.FormValue("ID"))

	character.UserID = getInt8FromString(r.FormValue("UserID"))

	character.Points = getInt8FromString(r.FormValue("Points"))

	character.Strength = getInt8FromString(r.FormValue("Strength"))

	character.Endurance = getInt8FromString(r.FormValue("Endurance"))

	character.Armor = getInt8FromString(r.FormValue("Armor"))

	character.FirePower = getInt8FromString(r.FormValue("FirePower"))

	character.Gold = getInt8FromString(r.FormValue("Gold"))

	character.Silver = getInt8FromString(r.FormValue("Silver"))

	character.Bronze = getInt8FromString(r.FormValue("Bronze"))

	character.MaxHp = getInt8FromString(r.FormValue("MaxHp"))

	character.MaxMp = getInt8FromString(r.FormValue("MaxMp"))

	character.Hp = getInt8FromString(r.FormValue("Hp"))

	character.Mp = getInt8FromString(r.FormValue("Mp"))

	character.Exp = getInt8FromString(r.FormValue("Exp"))

	character.Items = r.FormValue("Items")

	time := time.Now()

	if character.ID == 0 {
		character.Register = time
	} else {
		character.Updated = time
	}

	if len(character.Histories) <= 0 || r.FormValue("Histories") != character.Histories {
		character.Histories = r.FormValue("Histories")
	}

	if len(character.Expertises) <= 0 || r.FormValue("Expertises") != character.Expertises {
		character.Expertises = r.FormValue("Expertises")
	}

	if len(character.Advantages) <= 0 || r.FormValue("Advantages") != character.Advantages {
		character.Advantages = r.FormValue("Advantages")
	}

	if len(character.Disadvantages) <= 0 || r.FormValue("Disadvantages") != character.Disadvantages {
		character.Disadvantages = r.FormValue("Disadvantages")
	}

	if len(character.Spells) <= 0 || r.FormValue("Spells") != character.Spells {
		character.Spells = r.FormValue("Spells")
	}

	if character.ID == 0 {
		db.CreateCharacter(character)
	} else {
		db.UpdateCharacter(character)
	}

	characterSheet, err := ioutil.ReadFile("./pages/character_sheet.html")
	if err != nil {
		fmt.Print(err)
	}

	fullData := map[string]interface{}{
		"NavigationBar": template.HTML(navigationBarHTML),
		"Page":          template.HTML(characterSheet),
	}
	w.WriteHeader(http.StatusCreated)

	render(w, r, homepageTpl, "homepage_view", fullData)
}

func getInt8FromString(input string) int8 {
	if input == "" {
		return 0
	}
	val, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		panic(err)
	}
	return int8(val)
}
