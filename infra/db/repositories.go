package db

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

//Userr data to login
type Userr struct {
	ID         int8
	Name       string `sql:",unique"`
	Email      string `sql:",unique"`
	Password   string
	Access     int8 `sql:"default:0"`
	Characters []*Character
}

//Character of RPG
type Character struct {
	UserID        int8
	ID            int8
	Name          string
	Description   string
	Points        int8
	Strength      int8
	Ability       int8
	Endurance     int8
	Armor         int8
	FirePower     int8
	Gold          int8
	Silver        int8
	Bronze        int8
	MaxHp         int8
	MaxMp         int8
	Hp            int8
	Mp            int8
	Exp           int8
	Items         map[string]string `pg:",hstore"`
	Register      time.Time         `sql:"default:now()"`
	Updated       time.Time
	Histories     []*History
	Race          *Race
	Spells        []*Spell        `pg:"many2many:character_spells"`
	Expertises    []*Expertise    `pg:"many2many:character_expertises"`
	Advantages    []*Advantage    `pg:"many2many:character_advantages"`
	Disadvantages []*Disadvantage `pg:"many2many:character_disadvantages"`
}

//Race of RPG
type Race struct {
	ID                int8
	Name              string
	Description       string
	ChangedAttributes map[string]int8 `pg:",hstore"`
	Expertises        []*Expertise    `pg:"many2many:race_expertises"`
	Advantages        []*Advantage    `pg:"many2many:race_advantages"`
	Disadvantages     []*Disadvantage `pg:"many2many:race_disadvantages"`
}

//Advantage of RPG
type Advantage struct {
	ID                int8
	Name              string
	Description       string
	ChangedAttributes map[string]int8 `pg:",hstore"`
	Characters        []*Character    `pg:"many2many:character_advantages"`
	Races             []*Race         `pg:"many2many:race_advantages"`
}

//Disadvantage of RPG
type Disadvantage struct {
	ID                int8
	Name              string
	Description       string
	ChangedAttributes map[string]int8 `pg:",hstore"`
	Characters        []*Character    `pg:"many2many:character_disadvantages"`
	Races             []*Race         `pg:"many2many:race_disadvantages"`
}

//Expertise of RPG
type Expertise struct {
	ID          int8
	Name        string
	Description string
	Characters  []*Character `pg:"many2many:character_expertises"`
	Races       []*Race      `pg:"many2many:race_expertises"`
}

//Spell of RPG
type Spell struct {
	ID          int8
	Name        string
	Description string
	Characters  []*Character `pg:"many2many:character_spells"`
}

//History of Characters
type History struct {
	ID       int8
	Text     string
	Register time.Time `sql:"default:now()"`
}

//Report of admin(like histories i think)
type Report struct {
	ID       int8
	Title    string
	Text     string
	Register time.Time `sql:"default:now()"`
}

//
//*****************************
// MANY TO MANY RELATIONSHIP **
//*****************************
//

//CharacterSpell is many to many relationship
type CharacterSpell struct {
	CharacterID int `sql:",pk"` // pk tag is used to mark field as primary key
	Character   *Character
	SpellID     int `sql:",pk"`
	Spell       *Spell
}

//CharacterExpertise is many to many relationship
type CharacterExpertise struct {
	CharacterID int `sql:",pk"` // pk tag is used to mark field as primary key
	Character   *Character
	ExpertiseID int `sql:",pk"`
	Expertise   *Expertise
}

//CharacterAdvantage is many to many relationship
type CharacterAdvantage struct {
	CharacterID int `sql:",pk"` // pk tag is used to mark field as primary key
	Character   *Character
	AdvantageID int `sql:",pk"`
	Advantage   *Advantage
}

//CharacterDisadvantage is many to many relationship
type CharacterDisadvantage struct {
	CharacterID    int `sql:",pk"` // pk tag is used to mark field as primary key
	Character      *Character
	DisadvantageID int `sql:",pk"`
	Disadvantage   *Disadvantage
}

//RaceExpertise is many to many relationship
type RaceExpertise struct {
	RaceID      int `sql:",pk"` // pk tag is used to mark field as primary key
	Race        *Race
	ExpertiseID int `sql:",pk"`
	Expertise   *Expertise
}

//RaceAdvantage is many to many relationship
type RaceAdvantage struct {
	RaceID      int `sql:",pk"` // pk tag is used to mark field as primary key
	Race        *Race
	AdvantageID int `sql:",pk"`
	Advantage   *Advantage
}

//RaceDisadvantage is many to many relationship
type RaceDisadvantage struct {
	RaceID         int `sql:",pk"` // pk tag is used to mark field as primary key
	Race           *Race
	DisadvantageID int `sql:",pk"`
	Disadvantage   *Disadvantage
}

//
//*****************************
// Create Schemas
//*****************************
//

func createSchemas(db *pg.DB) error {
	for _, model := range []interface{}{
		(*Userr)(nil),
		(*Character)(nil),
		(*Race)(nil),
		(*Advantage)(nil),
		(*Disadvantage)(nil),
		(*Expertise)(nil),
		(*Spell)(nil),
		(*History)(nil),
		(*Report)(nil),
		(*CharacterAdvantage)(nil),
		(*CharacterDisadvantage)(nil),
		(*CharacterExpertise)(nil),
		(*CharacterSpell)(nil),
		(*RaceAdvantage)(nil),
		(*RaceDisadvantage)(nil),
		(*RaceExpertise)(nil),
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:          true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateSchema(db *pg.DB) error {
	for _, model := range []interface{}{
		(*Userr)(nil),
	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

//TestConnection .-.
func TestConnection() {
	db := GetConnect()
	defer db.Close()

	err := createSchemas(db)
	if err != nil {
		panic(err)
	}

	user := &Userr{
		Name:     "admin",
		Email:    "wisneymaciel2@hotmail.com",
		Password: "1234",
	}

	character := &Character{
		Name:   "Porcos",
		UserID: 1,
	}

	user.Characters = []*Character{character}

	err = db.Insert(user)
	if err != nil {
		panic(err)
	}

	err = db.Insert(character)
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	userStored := &Userr{ID: user.ID}
	err = db.Select(userStored)
	if err != nil {
		panic(err)
	}

	// Select user by primary key.
	charStored := &Character{ID: character.ID}
	err = db.Select(charStored)
	if err != nil {
		panic(err)
	}

	// Select all users.
	var users []Userr
	err = db.Model(&users).Where("name = 'admin'").Select()
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
	fmt.Println(userStored.Characters[0])
	fmt.Println(users)
	fmt.Println(charStored)
}

//ExistEmail return if email exist in db
func ExistEmail(email string) bool {
	db := GetConnect()
	defer db.Close()

	userr := db.Model(new(Userr)).Where("email = ?", email).Select()
	return userr == nil
}
