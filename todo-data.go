package main

import (
	"fmt"
)

func (t todoTable) put() error {

	updateQ := fmt.Sprintf("update shard_1.todo set title = '%v', status = '%v' where id = %v", t.Title, t.Status, t.ID)

	_, err := db.Query(updateQ)
	if err != nil {
		return err
	}

	return nil
}
func put(t todoTable) error {

	updateQ := fmt.Sprintf("update shard_1.todo set title = '%v', status = '%v' where id = %v", t.Title, t.Status, t.ID)

	_, err := db.Query(updateQ)
	if err != nil {
		return err
	}

	return nil
}

func get(id int) (todoTable, error) {

	selectQ := fmt.Sprintf("select * from shard_1.todo where id = %v LIMIT 1", id)

	rows, err := db.Query(selectQ)
	if err != nil {
		return todoTable{}, err
	}

	var t todoTable
	for rows.Next() {

		err = rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			return todoTable{}, err
		}
		break // as we are expecting only one record
	}

	return t, nil
}
