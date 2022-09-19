package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/vishwaszadte/go-gin-todo/routes"
)

func main() {

	r := gin.Default()

	v1 := r.Group("/api")
	routes.RegisterTodoRoutes(v1.Group("/"))

	fmt.Println("Starting the server")
	r.Run()
}
