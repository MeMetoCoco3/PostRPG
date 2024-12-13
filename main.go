package main

import (
	"PostRPG/internal/database"
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	ctx := context.Background()
	DealWithError(err)
	DB_URL := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", DB_URL)
	DealWithError(err)
	dbQueries := database.New(db)
	dbQueries.DeleteAllCharacters(ctx)
	dbQueries.DeleteAllSkills(ctx)
	dbQueries.DeleteAllWeapons(ctx)
	// Create a bunch of shit and test, then create assign weappon or skill to character,
	// then see if that works correctly, maybe create querie get  character skills/ weapons
	// when all is working
	// Create battles with characters.

	char1 := NewCharacter("Vidal El Rey", WARRIOR, "$")
	char2 := NewCharacter("Katerina", WIZZARD, "A")
	char1, _ = char1.UploadToDb(dbQueries)
	char2, _ = char2.UploadToDb(dbQueries)
	// func NewWeapon(name string, description string, damage int, reach int, role Role) *Weapon
	weapon1 := NewWeapon("Quizzizle", "Most powerfull kickass drumstick in the world", 10, 2, WARRIOR)
	weapon1, _ = weapon1.UploadToDb(dbQueries)
	// func NewSkill(name string, description string, damage int, reach int, role Role, cost int, coin string) *Skill {
	skill1 := NewSkill("Firebunga", "Tremendous projectile of kowabunga", 12, 4, WIZZARD, 12, "MANA")
	skill1, _ = skill1.UploadToDb(dbQueries)
}

func DealWithError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
