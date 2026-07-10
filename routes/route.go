package routes

import (
	"database/sql"

	"github.com/julienschmidt/httprouter"

	"todolist-api/controllers"
	"todolist-api/repositories"
	"todolist-api/services"
)

func Init(router *httprouter.Router, db *sql.DB) {
	todoRepo := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepo)
	todoController := controllers.NewTodoController(todoService)

	router.GET("/api/v1/todos", todoController.GetAll)
	router.GET("/api/v1/todos/:id", todoController.GetByID)
	router.POST("/api/v1/todos", todoController.Create)
	router.PUT("/api/v1/todos/:id", todoController.Update)
	router.PATCH("/api/v1/todos/:id/complete", todoController.MarkComplete)
	router.PATCH("/api/v1/todos/:id/uncomplete", todoController.MarkIncomplete)
	router.DELETE("/api/v1/todos/:id", todoController.Delete)
}