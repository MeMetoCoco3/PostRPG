package main

/*
import (
	"PostRPG/internal/database"
	"context"
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)


func main() {
	fmt.Println("MAIN")
	err := godotenv.Load()
	ctx := context.Background()
	DealWithError(err, "Error loading Env variables")
	DB_URL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", DB_URL)
	DealWithError(err, "Error starting database conection")
	dbQueries := database.New(db)
	dbQueries.DeleteAllCharacters(ctx)
	dbQueries.DeleteAllSkills(ctx)
	dbQueries.DeleteAllWeapons(ctx)
	err = dbQueries.SetAllNull(ctx)
	fmt.Printf("SetAllNull: %s\n", err)

	// func NewWeapon(name string, description string, damage int, reach int, role Role) *Weapon
	weapon1 := NewWeapon("Quizzizle", "Most powerfull kickass drumstick in the world", 10, 2, WARRIOR)
	weapon1, _ = weapon1.UploadToDb(dbQueries)
	// func NewSkill(name string, description string, damage int, reach int, role Role, cost int, coin string) *Skill {
	skill1 := NewSkill("Firebunga", "Tremendous projectile of kowabunga", 12, 4, WIZZARD, 12, "MANA")

	skill2 := NewSkill("Petrolox", "Ostion brutal", 20, 1, WARRIOR, 23, "STAMINA")
	skill3 := NewSkill("SomeBullshit", "Bla", 1, 2, WIZZARD, 1, "jajaj")
	skill1, err = skill1.UploadToDb(dbQueries)
	//fmt.Println(err)
	skill2, err = skill2.UploadToDb(dbQueries)
	//fmt.Println(err)
	skill3, err = skill3.UploadToDb(dbQueries)
	//fmt.Println(err)
	char1 := NewCharacter("Vidal El Rey", WARRIOR, "$")
	char2 := NewCharacter("Katerina", WIZZARD, "A")
	char1, err = char1.UploadToDb(dbQueries)
	//fmt.Println(err)
	char2, err = char2.UploadToDb(dbQueries)
	//fmt.Println(err)
	char1.GetWeapon(dbQueries, weapon1)
	fmt.Println(char1.Weapon)
	err = char2.GetSkill(dbQueries, skill1)
	fmt.Println(err)
	fmt.Println(char2.Skill)
	//fmt.Println(err)
	//fmt.Println(char2.Skill)
	//fmt.Println(weapon1)
	//fmt.Println(char1.Weapon)
	err = char1.Attack(dbQueries, char2, WEAPON)

	positions := []Position{
		{X: 1, Y: 3},
		{X: 3, Y: 5},
		{X: 5, Y: 3},
		{X: 1, Y: 2},
		{X: 0, Y: 0},
	}

	for _, position := range positions {
		err = char1.Move(dbQueries, position)
		_, _ = dbQueries.GetCharacterPosition(ctx, ToNullString(char1.Icon))
		// fmt.Printf("Character is in position (%d,%d)\n", terrainInfo.X, terrainInfo.Y)
	}

	_ = char1.Move(dbQueries, Position{X: 1, Y: 2})
	_ = char2.Move(dbQueries, Position{X: 1, Y: 3})
	fmt.Println(char2.Weapon)

	err = char1.Attack(dbQueries, char2, WEAPON)
	err = char2.Attack(dbQueries, char1, SKILL)
	if err != nil {
		fmt.Println(err)
	}
	err = char2.Move(dbQueries, Position{X: 1, Y: 2})
	if err != nil {
		fmt.Println(err)
	}
	// TODO: SetWeapons null
	// TODO: SetSkills null
	dbQueries.SetAllNull(ctx)
}
*/
