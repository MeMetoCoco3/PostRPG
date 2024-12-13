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

	// Create a bunch of shit and test, then create assign weappon or skill to character,
	// then see if that works correctly, maybe create querie get  character skills/ weapons
	// when all is working
	// Create battles with characters.

	char1 := NewCharacter("Vidal El Rey", WARRIOR, "$")
	char2 := NewCharacter("Katerina", WIZZARD, "A")
	char1.UploadToDb(dbQueries)
	char2.UploadToDb(dbQueries)

	//skill1 := NewSkill

}

func DealWithError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return
}
