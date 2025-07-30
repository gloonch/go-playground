package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

var todos = []Todo{
	{ID: 1, Text: "First task"},
	{ID: 2, Text: "Second task"},
	{ID: 3, Text: "Third task"},
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = len(todos) + 1
	todos = append(todos, todo)
	json.NewEncoder(w).Encode(todo)
}

func main() {
	http.HandleFunc("/todos", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method == http.MethodGet {
			getTodos(writer, request)
		} else if request.Method == http.MethodPost {
			addTodo(writer, request)
		} else {
			http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
