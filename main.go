package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var db *sql.DB

func initDb() error {
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

func main() {
	err := initDb()
	if err != nil {
		log.Fatalf("db init failed: %s", err.Error())
		return
	}

	port := "8080"

	router := mux.NewRouter()
	router.HandleFunc("/todo/{id}", editTodoHandler).Methods("PUT")

	log.Printf("app running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
