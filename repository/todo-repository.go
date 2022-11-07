package repository

import (
	"example/todolist-api/models"
	"log"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo models.Todo) models.Todo
	GetAllTodos(status string) []models.Todo
	GetTodo(id int) models.Todo
	//UpdateTodo(todo models.Todo) models.Todo
}

type todoConnection struct {
	connection *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoConnection{
		connection: db,
	}
}

func (db *todoConnection) CreateTodo(todo models.Todo) models.Todo {
	log.Println("Inside CreateTodo repo", todo)
	db.connection.Save(&todo)
	return todo
}

func (db *todoConnection) GetAllTodos(status string) []models.Todo {
	var todos []models.Todo
	if status == "Incomplete" || status == "complete" {
		db.connection.Where("status = ?", status).Find(&todos)
	} else if status == "" {
		db.connection.Find(&todos)
	}

	return todos
}

func (db *todoConnection) GetTodo(id int) models.Todo {
	var todo models.Todo
	db.connection.Where("id = ?", id).Find(&todo)
	return todo
}

// func (db *todoConnection) UpdateTodo(todo models.Todo) models.Todo {

// }
