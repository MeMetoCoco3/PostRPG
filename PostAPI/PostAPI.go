package PostAPI

import (
	"PostRPG/createdb"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

type Role int

const (
	WARRIOR Role = iota
	WIZZARD
	ARCHER
)

type Skill struct {
	payment     []int
	damage      int
	distance    int
	canBeLearnt Role
	name        string
	description string
	payWith     []string
}

type Weapon struct {
	damage      int
	reach       int
	name        string
	description string
}

type Character struct {
	ID        int    `json:"id"`
	Health    int    `json:"health"`
	Mana      int    `json:"mana"`
	Stamina   int    `json:"stamina"`
	Strength  int    `json:"strength"`
	Job       Role   `json:"role"`
	JobString string `json:"jobString"`
	Name      string `json:"name"`
	Skill     Skill  `json:"skill"`
	Weapon    Weapon `json:"weapon"`
	Icon      string `json:"icon"`
}

func NewCharacter(name string, job Role, icon string) *Character {
	var role string
	var health, mana, stamina, strength int
	switch job {
	case 0:
		role = "WARRIOR"
		health = 100
		mana = 0
		stamina = 4
		strength = 10
	case 1:
		role = "WIZZARD"
		health = 30
		mana = 30
		stamina = 2
		strength = 2
	case 2:
		role = "ARCHER"
		health = 50
		mana = 0
		stamina = 4
		strength = 6
	default:
		panic("Not correct job")
	}

	c := &Character{
		Health:    health,
		Mana:      mana,
		Stamina:   stamina,
		Strength:  strength,
		Job:       job,
		JobString: role,
		Name:      name,
		Icon:      icon,
	}
	return c
}

func UploadCharacterToDb(c *Character, db *sql.DB) error {
	jsonCharacter, err := json.Marshal(c)
	if err != nil {
		return errors.New("Not possible to jsonify this character!")
	}
	fmt.Println(string(jsonCharacter))
	_, err = db.Exec("INSERT INTO characters (details) VALUES ($1)", string(jsonCharacter))
	if err != nil {
		return errors.New("Error uploading character to the database")
	}
	return nil
}

func (c *Character) PrintStats() {
	fmt.Println("🌟 Character Sheet 🌟")
	fmt.Println("--------------------")
	fmt.Printf("Name:           %s\n", c.Name)
	fmt.Printf("Icon:           %s\n", c.Icon)
	fmt.Printf("Job:            %s\n", c.JobString)
	fmt.Println()
	fmt.Println("    📊 Stats 📊")
	fmt.Printf("Health:             %d\n", c.Health)
	fmt.Printf("Mana:               %d\n", c.Mana)
	fmt.Printf("Stamina:            %d\n", c.Stamina)
	fmt.Printf("Strength:           %d\n", c.Strength)
	fmt.Println()
	/*
		fmt.Println("🔮 Skill:")
		fmt.Printf("Name:       %s\n", skill.Name)
		fmt.Printf("Power:      %d\n", skill.Power)
		fmt.Println()
		fmt.Println("🗡️ Weapon:")
		fmt.Printf("Name:       %s\n", weapon.Name)
		fmt.Printf("Damage:     %d\n", weapon.Damage) fmt.Printf("Range:      %d\n", weapon.Range) fmt.Println("--------------------") */
}

func GetCharacterStats(i int) (Character, error) {
	db := createdb.StartConexion()
	result, err := db.Exec("SELECT * FROM characters WHERE id = $1", i)
	if err != nil {
		return Character{}, errors.New("Character Not Found")
	}
	fmt.Println(result)

	return Character{}, nil
}
