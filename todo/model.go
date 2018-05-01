package todo

import (
	"fmt"
)

const table = "public.todo"

var allowedStatuses = []string{"New", "In Progress", "Closed"}

// Todos is a list of todo.Todo structs
type Todos struct {
	TodoList []Todo `json:"todos"`
}

// Todo is a struct containing the ID of a todo, as well as, title and status
type Todo struct {
	ID int `json:"id"`
	CreateTodo
}

// CreateTodo is the expected payload for a create todo request
type CreateTodo struct {
	Title  string `json:"title"`
	Status string `json:"status"`
}

func get(id int) (Todo, error) {
	var t Todo

	selectQ := fmt.Sprintf("select * from %s where id = %v LIMIT 1", table, id)
	err := db.QueryRow(selectQ).Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		return Todo{}, err
	}

	return t, nil
}

func put(t Todo) error {
	updateQ := fmt.Sprintf("update %s set title = '%v', status = '%v' where id = %v", table, t.Title, t.Status, t.ID)

	_, err := db.Query(updateQ)
	if err != nil {
		return err
	}

	return nil
}

// isValid return appropriate message if the given todo struct is invalid, else returns blank string
func isValid(t CreateTodo) (invalidMsg string) {

	// blank fields check
	if t.Title == "" || t.Status == "" {
		return "Todo request is missing status or title"
	}

	// valid status check
	statusValid := false
	for _, s := range allowedStatuses {
		if t.Status == s {
			statusValid = true
			break
		}
	}
	if !statusValid {
		return "The provided status is not supported"
	}

	//
	return ""
}
