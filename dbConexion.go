package main

import (
	"PostRPG/internal/database"
	"context"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

func GetConexion() *database.Queries {
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
	dbQueries.SetAllNull(ctx)
	return dbQueries
}
