package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muly/todo-api/todo" // github.com/rackerlabs/GoCodingChallenge/todo
)

// Status := //TODO: add documentation
func Status(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println("Status Request Received")
	w.WriteHeader(200)
	fmt.Fprint(w, "OK\n")
}

func main() {
	err := todo.InitDb()
	if err != nil {
		log.Fatalf("db init failed: %s", err.Error())
		return
	}

	port := "8080"

	router := httprouter.New()
	router.GET("/", Status)
	router.POST("/todos", todo.Create)
	router.GET("/todos", todo.List)
	router.PUT("/todos", todo.Update)

	log.Printf("Starting server on port %s...\n", port)

	// Note: Make sure you have DB_USER, DB_PASSWORD and DB_NAME environment variables set.
	// We use them elsewhere
	log.Fatal(http.ListenAndServe(":"+port, router))
}
