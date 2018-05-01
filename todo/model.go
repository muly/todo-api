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

func getAll() ([]Todo, error) {
	var ts []Todo

	selectQ := fmt.Sprintf("select * from %s", table)
	rows, err := db.Query(selectQ)

	for rows.Next() {
		t := Todo{}
		rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			return []Todo{}, err
		}
		ts = append(ts, t)
	}

	return ts, nil
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

func post(t CreateTodo) (int, error) {
	insertQ := fmt.Sprintf(`INSERT INTO todo (title, status) VALUES ('%s', '%s') RETURNING id`, t.Title, t.Status)

	var todoID int
	// Insert and get back newly created todo ID
	if err := db.QueryRow(insertQ).Scan(&todoID); err != nil {
		return 0, err
	}

	return todoID, nil
}

func put(t Todo) error {
	updateQ := fmt.Sprintf("update %s set title = '%v', status = '%v' where id = %v", table, t.Title, t.Status, t.ID)

	_, err := db.Query(updateQ)
	if err != nil {
		return err
	}

	return nil
}
