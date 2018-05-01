package todo

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB //NOTE: move db connection code to its own common repo when there are more than one model

func InitDb() error {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := "disable"
	host := "localhost"
	port := "5433" //TODO: change this back to 5432

	dburl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbname, sslmode)

	var err error
	db, err = sql.Open("postgres", dburl)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	log.Println("DB connected")
	return nil
}
