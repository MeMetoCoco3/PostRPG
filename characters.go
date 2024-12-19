package main

import (
	"PostRPG/internal/database"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type Character struct {
	ID       uuid.UUID `json:"id"`
	Position Position
	Health   int     `json:"health"`
	Mana     int     `json:"mana"`
	Stamina  int     `json:"stamina"`
	Strength int     `json:"strength"`
	Role     Role    `json:"role"`
	Name     string  `json:"name"`
	Skill    *Skill  `json:"skill"`
	Weapon   *Weapon `json:"weapon"`
	Icon     string  `json:"icon"`
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
		Role:     job,
		Name:     name,
		Icon:     icon,
	}

	return c
}

func (c *Character) UploadToDb(db *database.Queries) (*Character, error) {
	ctx := context.Background()
	data, err := db.CreateNewCharacter(ctx, c.ToParams())
	err = DealWithError(err, "Database Upload")

	character := ParamsToCharacter(data)

	return character, err
}

func (c *Character) GetWeapon(db *database.Queries, w *Weapon) error {
	if c.Role != w.Role {
		return fmt.Errorf("This character cannot equip this weapon")
	}

	ctx := context.Background()
	err := db.AssignWeapon(ctx, database.AssignWeaponParams{
		ID:       c.ID,
		WeaponID: ToNullUUID(w.ID),
	})
	err = DealWithError(err, "Assign weapon to character")
	c.Weapon = w
	return err
}

func (c *Character) GetSkill(db *database.Queries, s *Skill) error {
	if c.Role != s.Role {
		return fmt.Errorf("This character cannot equip this weapon")
	}
	ctx := context.Background()
	err := db.AssignSkill(ctx, database.AssignSkillParams{
		ID:      c.ID,
		SkillID: ToNullUUID(s.ID),
	})
	err = DealWithError(err, "Asign skill to character")
	c.Skill = s
	return err
}

func (c *Character) Move(db *database.Queries, p Position) error {
	ctx := context.Background()
	result, err := db.CheckPosition(ctx, database.CheckPositionParams{
		X: int32(p.X),
		Y: int32(p.Y),
	})
	if err != nil {
		DealWithError(err, "(-) Error encountered: Checking the position while moving: %s")
		return err
	}
	if result.Character.Valid {
		return fmt.Errorf("(-) Error encountered: A character(%s) is already in position (%d,%d)", result.Character.String, p.X, p.Y)
	}

	err = db.SetNull(ctx, database.SetNullParams{X: int32(c.Position.X), Y: int32(c.Position.Y)})
	if err != nil {
		return err
	}
	err = db.MoveCharacter(ctx, database.MoveCharacterParams{
		Character: ToNullString(c.Icon),
		X:         int32(p.X),
		Y:         int32(p.Y),
	})

	DealWithError(err, "(-) Error encountered: Error moving the character to Postgres DB\n")
	c.Position = p
	return nil
}

func (c *Character) Attack(db *database.Queries, objective *Character, action Action) error {
	ctx := context.Background()
	newHealth := objective.Health
	switch action {
	case ATTACK:
		// TODO:CHECK REACH
		if c.Stamina < 1 {
			return fmt.Errorf("(-) Character does not have enough stamina.\n")
		}
		newHealth -= c.Strength
		c.Stamina--
	case WEAPON:
		if c.Weapon == nil {
			return fmt.Errorf("(-) Character does not have weapon.\n")
		}
		// TODO: CHECK REACH
		if c.Stamina < 2 {
			return fmt.Errorf("(-) Character does not have enough stamina.\n")
		}
		if c.Weapon == nil {
			return fmt.Errorf("(-) Character does not have a weapon.\n")
		}
		newHealth -= (c.Strength + c.Weapon.Damage)
		c.Stamina -= 2
	case SKILL:
		// TODO: CHECK REACH
		if c.Skill == nil {
			return fmt.Errorf("(-) Character does not have skill.\n")
		}
		switch c.Skill.Coin {
		case "MANA":
			if c.Skill.AmountToPay > c.Mana {
				return fmt.Errorf("Not enough mana. -%d, +%d.\n", c.Skill.AmountToPay, c.Mana)
			}
			c.Mana -= c.Skill.AmountToPay
		case "STAMINA":
			if c.Skill.AmountToPay > c.Stamina {
				return fmt.Errorf("Not enough stamina. -%d, +%d.\n", c.Skill.AmountToPay, c.Stamina)
			}
			c.Stamina -= c.Skill.AmountToPay
		default:
			return fmt.Errorf("Not correct coin for skill.\n")
		}
		if c.Skill == nil {
			return fmt.Errorf("Character does not have a skill.\n")
		}
		newHealth = newHealth - (c.Strength + c.Skill.Damage)
	default:
		return fmt.Errorf("Character.Attack method: This action is not allowed.\n")
	}

	err := db.SetHealth(ctx, database.SetHealthParams{
		ID:     objective.ID,
		Health: int32(newHealth),
	})
	c.Health = newHealth
	err = DealWithError(err, "Asign skill to character")

	return err
}

func (c *Character) ToParams() database.CreateNewCharacterParams {
	return database.CreateNewCharacterParams{
		Health:   int32(c.Health),
		Mana:     int32(c.Mana),
		Stamina:  int32(c.Stamina),
		Strength: int32(c.Strength),
		Job:      int32(c.Role),
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
		Role:     Role(data.Job),
		Name:     data.Name,
		Skill:    nil,
		Weapon:   nil,
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
