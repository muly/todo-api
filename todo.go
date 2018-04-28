package main

import ()

type todo struct { // {"title":"car service", "status": "open"}
	Title  string `json:"title,omitempty"`
	Status string `json:"status,omitempty"`
}

type todoTable struct {
	ID int `json:"id,omitempty"`
	todo
}
