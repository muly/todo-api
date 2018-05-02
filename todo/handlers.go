package todo

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Create will allow a user to create a new todo
// The supported body is {"title": "", "status": ""}
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// read body
	t := CreateTodo{}
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		msg := "invalid json data"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	if invalidMsg := isValid(t); invalidMsg != "" {
		msg := "invalid todo message: " + invalidMsg
		log.Printf("%s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	id, err := post(t)
	if err != nil {
		msg := "insert failed"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	newTodo, err := get(id)
	if err != nil {
		msg := "fetch failed"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newTodo)
	if err != nil {
		msg := "json encode failed"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

}

// List will provide a list of all current to-dos
func List(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todoList, err := getAll()
	if err != nil {
		msg := "getall failed:"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(Todos{TodoList: todoList})
	if err != nil {
		msg := "json encode failed"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
}

// Update will allow a user to update an existing todo
// The supported body is {"title": "", "status": ""}
func Update(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		msg := "id must be an integer"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusBadRequest)
	}

	// read body
	t := CreateTodo{}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		msg := "invalid json data"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// checks
	if invalidMsg := isValid(t); invalidMsg != "" {
		msg := "invalid todo message: " + invalidMsg
		log.Printf("%s", msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}
	rec, err := get(id)
	if err == sql.ErrNoRows {
		msg := "id doesn't exist in db"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusBadRequest)
		return
	} else if err != nil {
		msg := ""
		log.Printf("fetch (as part of verify) failed:%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}
	log.Println("existing record from database:", rec)

	// update todo
	err = put(Todo{ID: id, CreateTodo: t})
	if err != nil {
		msg := "update failed"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	rec, err = get(id)
	if err != nil {
		msg := "fetch updated rec failed"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(rec)
	if err != nil {
		msg := "json encode failed"
		log.Printf("%s:%s", msg, err.Error())
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

}
