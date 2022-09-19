package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vishwaszadte/go-gin-todo/controllers"
)

func RegisterTodoRoutes(router *gin.RouterGroup) {
	router.GET("/todos", controllers.GetAllTodos)
	router.POST("/todos", controllers.CreateTodo)
	router.GET("/todos/:id", controllers.GetTodo)
	router.PUT("/todos/:id", controllers.UpdateTodo)
	router.DELETE("todos/:id", controllers.DeleteTodo)
}
