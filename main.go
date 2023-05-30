package main

import (
	"example/togo-go/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/todos", controller.GetTodos)
	r.POST("/todos", controller.AddTodos)
	r.GET("/todos/:id", controller.GetTodo)
	r.PATCH("/todos/:id", controller.UpdateTodoStatus)
	r.GET("/", controller.GetTodo)

	r.Run("localhost:9091")

}
