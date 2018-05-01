package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// edittodoHandler updates one todo item
// /todo/{id}
func editTodoHandler(w http.ResponseWriter, r *http.Request) {
	// read id param
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Fatalf("id param must be number: %s", err.Error())
		http.Error(w, "id param must be number", http.StatusBadRequest)
		return
	}
	fmt.Println(id)

	// read body
	t := todo{}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		log.Fatalf("invalid json data: %s", err.Error())
		http.Error(w, "invalid json data", http.StatusBadRequest)
		return
	}

	// update todo

	rec := todoTable{ID: id, todo: t}

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
