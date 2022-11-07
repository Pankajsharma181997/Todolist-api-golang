package services

import (
	"example/todolist-api/dto"
	"example/todolist-api/models"
	"example/todolist-api/repository"
	"log"

	"github.com/mashingan/smapping"
)

type TodoService interface {
	CreateTodo(todo dto.CreateTodoDTO) models.Todo
	GetAllTodos(status string) []models.Todo
	GetTodo(id int) models.Todo
	//UpdateTodo(todoId uint64) models.Todo
}

type todoService struct {
	todoRepository repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return &todoService{
		todoRepository: todoRepo,
	}
}

func (service *todoService) CreateTodo(todo dto.CreateTodoDTO) models.Todo {
	todoToCreate := models.Todo{}
	log.Println("Inside CreateTodo service")

	err := smapping.FillStruct(&todoToCreate, smapping.MapFields(&todo))

	if err != nil {
		log.Fatalf("Failed map %v", err)
	}

	res := service.todoRepository.CreateTodo(todoToCreate)
	return res

}

func (service *todoService) GetAllTodos(status string) []models.Todo {
	return service.todoRepository.GetAllTodos(status)
}

func (service *todoService) GetTodo(id int) models.Todo {
	return service.todoRepository.GetTodo(id)
}
