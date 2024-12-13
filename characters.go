package main

import (
	"PostRPG/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
)

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

func (c *Character) UploadToDb(db *database.Queries) (*Character, error) {
	ctx := context.Background()
	data, err := db.CreateNewCharacter(ctx, c.ToParams())
	DealWithError(err)

	character := ParamsToCharacter(data)

	return character, nil
}

func (c *Character) ToParams() database.CreateNewCharacterParams {
	return database.CreateNewCharacterParams{
		Health:   int32(c.Health),
		Mana:     int32(c.Mana),
		Stamina:  int32(c.Stamina),
		Strength: int32(c.Strength),
		Job:      int32(c.Job),
		Name:     c.Name,
		Icon:     c.Icon,
	}
}

func ParamsToCharacter(data database.Character) *Character {
	return &Character{
		ID:       data.ID,
		Health:   int(data.Health),
		Mana:     int(data.Mana),
		Stamina:  int(data.Stamina),
		Strength: int(data.Strength),
		Job:      int(data.Job),
		Name:     data.Name,
		Skill:    uuid.Nil,
		Weapon:   uuid.Nil,
		Icon:     data.Icon,
	}
}

/*
//  From the time I got trolled for not using UNIQUE/NotNull constraints in my fucking sqlc queries.

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

	func ParamsToCharacter(data database.Character) *Character {
		return &Character{
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
	}
*/
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
