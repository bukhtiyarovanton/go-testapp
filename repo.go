package main

import (
	"fmt"
)

var todos []Todo // Todos storage

var currentId int

func init() {
	RepoCreateTodo(Todo{Name: "Todo 1"})
	RepoCreateTodo(Todo{Name: "Todo 2"})
	RepoCreateTodo(Todo{Name: "Todo 3"})
}

func RepoFindTodo(id int) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	return Todo{}
}

func RepoCreateTodo(t Todo) Todo {
	currentId += 1
	t.Id = currentId
	todos = append(todos, t)
	return t
}

func RepoRemoveTodo(id int) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
