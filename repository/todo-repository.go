package repository

import (
	"example/todolist-api/dto"
	"example/todolist-api/models"
	"log"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo models.Todo) models.Todo
	GetAllTodos(status string) []models.Todo
	GetTodo(id int) models.Todo
	UpdateTodoStatus(id int) models.Todo
	UpdateTodo(id int, updates dto.CreateTodoDTO) models.Todo
	DeleteTodo(id int)
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

func (db *todoConnection) UpdateTodoStatus(id int) models.Todo {
	var todo models.Todo
	db.connection.Where("id = ?", id).Find(&todo)

	if todo.Status == "complete" {
		db.connection.Model(&todo).Where("id = ?", id).Update("status", "Incomplete")
	} else if todo.Status == "Incomplete" {
		db.connection.Model(&todo).Where("id = ?", id).Update("status", "complete")
	}
	db.connection.Where("id = ?", id).Find(&todo)
	return todo

}

func (db *todoConnection) UpdateTodo(id int, updates dto.CreateTodoDTO) models.Todo {
	var todo models.Todo

	db.connection.Where("id = ?", id).Find(&todo)

	todo.Status = updates.Status
	todo.Text = updates.Text

	db.connection.Save(&todo)

	return todo
}

func (db *todoConnection) DeleteTodo(id int) {
	var todo models.Todo

	db.connection.Where("id = ?", id).Delete(&todo)

}
