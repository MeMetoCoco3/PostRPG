package main

import (
	"PostRPG/internal/database"
	_ "PostRPG/internal/database"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
)

type Role int

const (
	WARRIOR Role = iota
	WIZZARD
	ARCHER
)

type Skill struct {
	Id          uuid.UUID `json:"id"`
	AmountToPay int
	Coin        string
	Damage      int    `json:"damage"`
	Reach       int    `json:"Reach"`
	Role        Role   `json:"role"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Weapon struct {
	ID          uuid.UUID `json:"id"`
	Damage      int       `json:"damage"`
	Role        Role      `json:"Reach"`
	Reach       int       `json:"reach"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type Character struct {
	ID       uuid.UUID `json:"id"`
	Health   int       `json:"health"`
	Mana     int       `json:"mana"`
	Stamina  int       `json:"stamina"`
	Strength int       `json:"strength"`
	Job      int       `json:"role"`
	Name     string    `json:"name"`
	Skill    uuid.UUID `json:"skill"`
	Weapon   uuid.UUID `json:"weapon"`
	Icon     string    `json:"icon"`
}

func (c *Character) ToParams() database.CreateNewCharacterParams {
	return database.CreateNewCharacterParams{
		Health:   ToNullInt32(c.Health),
		Mana:     ToNullInt32(c.Mana),
		Stamina:  ToNullInt32(c.Stamina),
		Strength: ToNullInt32(c.Strength),
		Job:      ToNullInt32(c.Job),
		Name:     ToNullString(c.Name),
		Icon:     ToNullString(c.Icon),
	}
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func ToNullUUID(u uuid.UUID) uuid.NullUUID {
	return uuid.NullUUID{
		UUID:  u,
		Valid: u != uuid.Nil,
	}
}

func ToNullInt32(i int) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(i),
		Valid: i != 0,
	}
}

func (c *Character) UploadToDb(db *database.Queries) (Character, error) {
	ctx := context.Background()
	data, err := db.CreateNewCharacter(ctx, c.ToParams())

	DealWithError(err)

	character := Character{
		ID:       data.ID,
		Health:   int(data.Health.Int32),
		Mana:     int(data.Mana.Int32),
		Stamina:  int(data.Stamina.Int32),
		Strength: int(data.Strength.Int32),
		Job:      int(data.Job.Int32),
		Name:     data.Name.String,
		Skill:    uuid.Nil,
		Weapon:   uuid.Nil,
		Icon:     data.Icon.String,
	}

	return character, nil
}

func (s *Skill) UploadToDb(db *database.Queries) (Character, error) {
	ctx := context.Background()
	_, err := db.CreateNewSkill(ctx, database.CreateNewSkillParams{
		Damage:      ToNullInt32(s.Damage),
		Reach:       ToNullInt32(s.Reach),
		Coin:        ToNullString(s.Coin),
		AmountToPay: ToNullInt32(s.AmountToPay),
		Name:        ToNullString(s.Name),
		Description: ToNullString(s.Description),
		Role:        ToNullInt32(int(s.Role)),
	})

	DealWithError(err)
	return nil
}

func (w *Weapon) UploadToDb(db *database.Queries) error {
	rawBin, err := json.Marshal(w)
	DealWithError(err)

	ctx := context.Background()
	_, err = db.CreateNewWeapon(ctx, rawBin)
	DealWithError(err)
	return nil
}

func NewCharacter(name string, job Role, icon string) *Character {
	var health, mana, stamina, strength int
	switch job {
	// WARRIOR
	case 0:
		health = 100
		mana = 0
		stamina = 4
		strength = 10
	// WIZZARD
	case 1:
		health = 30
		mana = 30
		stamina = 2
		strength = 2
	// ARCHER
	case 2:
		health = 50
		mana = 0
		stamina = 4
		strength = 6
	default:
		panic("Not correct job")
	}

	c := &Character{
		Health:   health,
		Mana:     mana,
		Stamina:  stamina,
		Strength: strength,
		Job:      int(job),
		Name:     name,
		Icon:     icon,
	}
	return c
}

func NewWeapon(name string, description string, damage int, reach int, role Role) *Weapon {
	return &Weapon{
		Name:        name,
		Description: description,
		Damage:      damage,
		Reach:       reach,
		Role:        role,
	}
}

func NewSkill(name string, description string, damage int, reach int, role Role, cost []int, coin []string) *Skill {
	if len(cost) != len(coin) {
		os.Exit(1)
		return nil
	}

	payment := make(map[string]int)
	for i, val := range coin {
		payment[val] = cost[i]
	}

	return &Skill{
		Name:        name,
		Description: description,
		Damage:      damage,
		Payment:     payment,
		Role:        role,
	}
}
func (c *Character) PrintStats() {
	fmt.Println("🌟 Character Sheet 🌟")
	fmt.Println("--------------------")
	fmt.Printf("Name:           %s\n", c.Name)
	fmt.Printf("Icon:           %s\n", c.Icon)
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

// TODO
/*
func GetCharacterStats(i int, db *database.Queries) (Character, error) {
	result, err := db.Exec("SELECT * FROM characters WHERE id = $1", i)
	if err != nil {
		return Character{}, errors.New("Character Not Found")
	}
	fmt.Println(result)

	return Character{}, nil
}*/
