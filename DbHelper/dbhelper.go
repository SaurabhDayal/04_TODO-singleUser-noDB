package DbHelper

import (
	"01_TODO-SingleUser-noDB/models"
	"errors"
	"fmt"
)

func FindTodoById(id int) (*models.Todo, error) {
	for _, t := range models.Todos {
		if t.Id == id {
			return &t, nil
		}
	}
	return &models.Todo{}, errors.New(fmt.Sprintf("No value found with the %v id", id))
}

func FindAllTodos() ([]models.Todo, error) {
	return models.Todos, nil
}

func CreateNewTodo(todo *models.Todo) (models.Todo, error) {
	todo2 := models.Todo{todo.Id, todo.UserId, todo.Title, todo.Description}
	models.Todos = append(models.Todos, todo2)
	return todo2, nil
}

func UpdateTodoById(todo models.Todo) (models.Todo, error) {
	pos := pos(todo.Id, models.Todos)
	todo2 := models.Todo{todo.Id, todo.UserId, todo.Title, todo.Description}
	models.Todos[pos] = todo2
	return todo2, nil
}

func DeleteTodoById(id int) (*models.Todo, error) {
	var findTodo *models.Todo
	findTodo, err := FindTodoById(id)
	if err != nil {
		return &models.Todo{}, err
	}
	pos := pos(id, models.Todos)
	if pos < 0 {
		return &models.Todo{}, err
	}
	models.Todos = append(models.Todos[:pos], models.Todos[pos+1:]...)
	return findTodo, nil
}

func pos(id int, slice []models.Todo) int {
	for p, todo := range slice {
		if todo.Id == id {
			return p
		}
	}
	return -1
}
