package main

import (
	"PostRPG/PostAPI"
	"PostRPG/createdb"
	_ "fmt"
)

func main() {
	c := PostAPI.NewCharacter("Vidal", PostAPI.WARRIOR, "8")
	db := createdb.StartConexion()
	//createdb.CreateTypes(db)
	createdb.CreateTables(db)

	PostAPI.UploadCharacterToDb(c, db)
}
