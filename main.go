package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tsvillain/go-todo-server/controller"
	"github.com/tsvillain/go-todo-server/middleware"
)

func main() {
	router := gin.Default()
	router.GET("/", controller.GetAllTodos)
	router.GET("/:username", middleware.ValidateUser(), controller.GetAllTodosOfUser)
	router.GET("/todo/:id", middleware.ValidateId(), controller.GetTodoById)
	router.POST("/todo", controller.AddTodo)
	router.DELETE("/todo/:id/:username", middleware.ValidateId(), middleware.ValidateUser(), controller.DeleteTodo)
	router.PUT("/todo", controller.UpdateTodo)
	router.Run("localhost:3000")
}
