package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vishwaszadte/go-gin-todo/models"
)

func GetAllTodos(c *gin.Context) {
	var todos []models.Todo
	err := models.GetTodos(&todos)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todos)
	}
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	c.BindJSON(&todo)
	err := models.CreateTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusCreated, todo)
	}
}

func GetTodo(c *gin.Context) {
	id := c.Params.ByName("id")

	var todo models.Todo

	c.BindJSON(&todo)

	err := models.GetTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, todo)
	}

}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")

	err := models.GetTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.BindJSON(&todo)
	err = models.UpdateTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteTodo(c *gin.Context) {
	var todo models.Todo
	id := c.Params.ByName("id")
	c.BindJSON(&todo)

	err := models.DeleteTodo(&todo, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id: " + id: "deleted",
		})
	}
}
