package todo

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB //NOTE: move db connection code to its own common repo when there are more than one model

// InitDb initializes the package level database connection object
func InitDb() error {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	var err error
	db, err = sql.Open("postgres", dbinfo)
	if err != nil {
		log.Println("sql.Open failed")
		return err
	}

	err = db.Ping()
	if err != nil {
		log.Println("db.Ping failed")
		return err
	}

	log.Println("DB connected")
	return nil
}
