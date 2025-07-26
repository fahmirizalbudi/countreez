package configs

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

var (
	DB  *sql.DB
	err error
)

func GetPsqlConnection() {
	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=disable`,
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGDATABASE"),
	)

	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("PostgreSQL connection established successfully.")
}
