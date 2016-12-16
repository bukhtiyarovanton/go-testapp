package main

import (
	"github.com/bukhtiyarovanton/go-testapp/models"
)

/*
func init() {
	RepoCreateTodo(models.Todo{Name: "Todo 1"})
	RepoCreateTodo(models.Todo{Name: "Todo 2"})
	RepoCreateTodo(models.Todo{Name: "Todo 3"})
}
*/

func TodoRepoGetAllForUser(userID uint) []models.Todo {
	var todos []models.Todo
	var user models.User

	db.First(&user, userID).Related(&todos)
	return todos
}

func TodoRepoFindForUser(userID uint, id uint) models.Todo {
	var todo models.Todo
	var user models.User
	//db.Where("id = ? AND user_id = ?", id, userID).First(&todo, id)
	db.First(&user, userID).First(&todo, id)
	return todo
}

func TodoRepoCreateForUser(userID uint, t models.Todo) models.Todo {
	t.UserID = userID
	db.Create(&t)
	return t
}

func TodoRepoDeleteForUser(userID uint, id uint) models.Todo {

	todoToDelete := models.Todo{
		UserID: userID,
		Model: models.Model{
			ID: id,
		},
	}

	db.Find(&todoToDelete).Delete(&todoToDelete)
	return todoToDelete
}

func UserRepoFindByName(name string) (models.User, error) {
	var user models.User
	db.Where("name = ?", name).Find(&user)
	return user, nil
}
