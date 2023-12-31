package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dbHost := os.Getenv("HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("USER")
	dbPort := os.Getenv("PORT")
	dbPass := os.Getenv("PASS")

	connectionString := fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
