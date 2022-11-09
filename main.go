package main

import (
	"example/todolist-api/connectors"
	"example/todolist-api/controllers"
	"example/todolist-api/repository"
	"example/todolist-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func init() {
	connectors.ConnectToDB()
}

var (
	db             *gorm.DB                   = connectors.ConnectToDB()
	todoRepository repository.TodoRepository  = repository.NewTodoRepository(db)
	todoService    services.TodoService       = services.NewTodoService(todoRepository)
	todoController controllers.TodoController = controllers.NewTodoService(todoService)
)

func main() {
	defer connectors.CloseDBConnection(db)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to todo List Api",
		})
	})

	todoRoutes := r.Group("api/v1")
	{
		todoRoutes.POST("/create", todoController.CreateTodo)
		todoRoutes.GET("/todos", todoController.GetAllTodos)
		todoRoutes.GET("/todos/:id", todoController.GetTodo)
		todoRoutes.PATCH("/todos/:id", todoController.UpdateTodoStatus)
		todoRoutes.PUT("/todos/:id", todoController.UpdateTodo)
		todoRoutes.DELETE("/todos/:id", todoController.DeleteTodo)
	}
	r.Run()
}
