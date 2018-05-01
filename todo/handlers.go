package todo

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Create will allow a user to create a new todo
// The supported body is {"title": "", "status": ""}
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(err.Error())
	}
	var todo CreateTodo

	json.NewDecoder(r.Body).Decode(&todo) // TODO: handle error

	// checks
	if invalidMsg := isValid(todo); invalidMsg != "" {
		http.Error(w, invalidMsg, http.StatusBadRequest)
		return
	}

	insertStmt := fmt.Sprintf(`INSERT INTO todo (title, status) VALUES ('%s', '%s') RETURNING id`, todo.Title, todo.Status)

	var todoID int

	// Insert and get back newly created todo ID
	if err := db.QueryRow(insertStmt).Scan(&todoID); err != nil {
		fmt.Printf("Failed to save to db: %s", err.Error())
	}

	fmt.Printf("Todo Created -- ID: %d\n", todoID)

	newTodo := Todo{}
	db.QueryRow("SELECT id, title, status FROM todo WHERE id=$1", todoID).Scan(&newTodo.ID, &newTodo.Title, &newTodo.Status)

	jsonResp, _ := json.Marshal(newTodo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, string(jsonResp))
}

// List will provide a list of all current to-dos
func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	dbUser := os.Getenv("DB_USER")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dbinfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		fmt.Println(err.Error())
	}

	todoList := []Todo{}

	rows, err := db.Query("SELECT id, title, status FROM todo")
	defer rows.Close()

	for rows.Next() {
		todo := Todo{}
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status); err != nil {
			w.WriteHeader(500)
			fmt.Fprintf(w, "Failed to build todo list")
		}

		todoList = append(todoList, todo)
	}

	jsonResp, _ := json.Marshal(Todos{TodoList: todoList})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	fmt.Fprintf(w, string(jsonResp))
}

// Update will allow a user to update an existing todo
// /todos?id=
// The supported body is {"title": "", "status": ""}
func Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	vars, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var id int
	if val, exists := vars["id"]; exists {
		id, err = strconv.Atoi(val[0])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	// read body
	t := CreateTodo{}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Fatalf("invalid json data: %s", err.Error())
		http.Error(w, "invalid json data", http.StatusBadRequest)
		return
	}

	// checks
	if invalidMsg := isValid(t); invalidMsg != "" {
		http.Error(w, invalidMsg, http.StatusBadRequest)
		return
	}

	// update todo
	rec := Todo{ID: id, CreateTodo: t}

	fmt.Println(rec)
	err = put(rec)
	if err != nil {
		log.Fatalf("update failed: %s", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	rec, err = get(rec.ID)
	if err != nil {
		log.Fatalf("fetch failed: %s", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		log.Fatalf("json encode failed: %s", err.Error())
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

}
