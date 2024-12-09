package createdb

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	_ "os"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "evildead20"
	dbname = "postrpg"
)

func StartConexion() *sql.DB {
	//password := os.Getenv("POSTPASS")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable",
		host, port, user, dbname)
	// Open validates the arguments
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	// Ping creates the conexion
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db

}

func CreateTypes(db *sql.DB) {
	fmt.Println("Creating types")
	queries := []string{
		`CREATE TYPE skill AS (
			payment INTEGER[],
			damage INTEGER,
			distance INTEGER,
			role INT,
			name TEXT,
			description TEXT,
			pay_with TEXT[]
		)`,
		`CREATE TYPE weapon AS (
			damage INTEGER,
			reach INTEGER,
			name TEXT,
			description TEXT
		)`,
		`CREATE TYPE characters AS (
			health INTEGER,
			mana INTEGER,
			stamina INTEGER,
			strength INTEGER,
			role int,
			name TEXT,
			skill skill,
			weapon weapon,
			icon VARCHAR(1)
		)`,
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("Types Created")
}

func CreateTables(db *sql.DB) {
	fmt.Println("Creating tables")
	queries := []string{
		`CREATE TABLE IF NOT EXISTS skills (
    id SERIAL PRIMARY KEY,
    details skill NOT NULL, 
	
	CONSTRAINT unique_name
    UNIQUE ((details).name),

	CONSTRAINT no_null_fields
    CHECK (
        (details).payment IS NOT NULL AND
		(details).damage IS NOT NULL AND
        (details).distance IS NOT NULL AND
        (details).can_be_learnt IS NOT NULL AND
        (details).name IS NOT NULL AND
        (details).description IS NOT NULL AND
        (details).pay_with IS NOT NULL
    )
    );`,
		/*,
				`CREATE TABLE IF NOT EXISTS weapons (
		    id SERIAL PRIMARY KEY,
		    details weapon NOT NULL,
		    UNIQUE (details.name),
		    CHECK(
		      details.damage IS NOT NULL AND
		      details.reach IS NOT NULL AND
		      details.name IS NOT NULL AND
		      details.description IS NOT NULL
		    )
		    );`,
				`CREATE TABLE IF NOT EXISTS characters (
		    id SERIAL PRIMARY KEY,
		    details character NOT NULL,
		    UNIQUE (details.name),
		    CHECK (
		      details.health IS NOT NULL AND
		      details.mana IS NOT NULL AND
		      details.stamina IS NOT NULL AND
		      details.strength IS NOT NULL AND
		      details.role IS NOT NULL AND
		      details.name IS NOT NULL AND
		      details.icon IS NOT NULL
		    )
		    );`,*/
	}

	for _, query := range queries {
		if _, err := db.Exec(query); err != nil {
			log.Fatalln(err)
		}
	}
	fmt.Println("Tables created")
	fmt.Println("Creating terrain")
	// Matrix. Unnest will return a set with duplicates of the values of
	// a nested matrix.
	query := `CREATE TABLE terrains {
    id SERIAL PRIMARY KEY,
    matrix INTEGER[][] CHECK{
      matrix IS NOT NULL AND
      array_length(matrix, 1) > 10 AND
      array_length(matrix, 2) > 10 AND
      NOT EXISTS (
      SELECT 1 FROM unnest(matrix) AS element
      WHERE element > 2 OR element < 0
      )
    }
  };`
	if _, err := db.Exec(query); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Terrain Created")
}
