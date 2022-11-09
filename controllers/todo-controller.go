package controllers

import (
	"example/todolist-api/dto"
	"example/todolist-api/helper"
	"example/todolist-api/models"
	"example/todolist-api/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoController interface {
	CreateTodo(ctx *gin.Context)
	GetAllTodos(ctx *gin.Context)
	GetTodo(ctx *gin.Context)
	UpdateTodoStatus(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}

type todoController struct {
	todoService services.TodoService
}

func NewTodoService(todoService services.TodoService) TodoController {
	return &todoController{
		todoService: todoService,
	}
}

func (c *todoController) CreateTodo(ctx *gin.Context) {
	var createDTO dto.CreateTodoDTO

	errDTO := ctx.ShouldBind(&createDTO)

	if errDTO != nil {
		response := helper.BulldErrorResponse("Failed to Process the request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	createdTodo := c.todoService.CreateTodo(createDTO)
	response := helper.BuildResponse(true, "OK!", createdTodo)
	ctx.JSON(http.StatusCreated, response)
}

func (c *todoController) GetAllTodos(ctx *gin.Context) {
	var status string
	status, isStatusInQueryParams := ctx.GetQuery("status")
	var allTodos []models.Todo

	var todosArr []dto.UpdateTodoDTO

	if isStatusInQueryParams && (status != "Incomplete" && status != "complete") {
		response := helper.BulldErrorResponse("Failed to Process the request", "Invalid status", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	errDTO := ctx.ShouldBind(todosArr)

	if errDTO != nil {
		response := helper.BulldErrorResponse("Failed to Process the request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if isStatusInQueryParams {
		allTodos = c.todoService.GetAllTodos(status)
	} else {
		allTodos = c.todoService.GetAllTodos("")
	}

	response := helper.BuildResponse(true, "OK!", allTodos)
	ctx.JSON(http.StatusOK, response)

}

func (c *todoController) GetTodo(ctx *gin.Context) {
	// var id uint64
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Fatal("Can't convert string to int")
	}

	todo := c.todoService.GetTodo(id)

	if todo.Text == "" {
		response := helper.BulldErrorResponse("Failed to Process the request", "No Todo available", helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
	} else {
		response := helper.BuildResponse(true, "OK!", todo)
		ctx.JSON(http.StatusOK, response)
	}

}

func (c *todoController) UpdateTodoStatus(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Fatal("Can't convert string to int")
	}

	todo := c.todoService.UpdateTodoStatus(id)
	response := helper.BuildResponse(true, "OK!", todo)
	ctx.JSON(http.StatusOK, response)
}

func (c *todoController) UpdateTodo(ctx *gin.Context) {

	var createDTO dto.CreateTodoDTO

	errDTO := ctx.ShouldBind(&createDTO)

	if errDTO != nil {
		response := helper.BulldErrorResponse("Failed to Process the request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Fatal("Can't convert string to int")
	}

	todo := c.todoService.UpdateTodo(id, createDTO)
	response := helper.BuildResponse(true, "OK!", todo)
	ctx.JSON(http.StatusOK, response)

}

func (c *todoController) DeleteTodo(ctx *gin.Context) {
	idParam := ctx.Param("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		log.Fatal("Can't convert string to int")
	}

	c.todoService.DeleteTodo(id)
	response := helper.BuildResponse(true, "Deleted element", helper.EmptyObj{})
	ctx.JSON(http.StatusOK, response)
}
