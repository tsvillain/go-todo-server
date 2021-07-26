package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Enum for Priority
type priority int

const (
	Low priority = iota
	Medium
	High
)

// Todo DataStructure
type todo struct{
	Id int `json:"id" binding:"required"`
	Task string `json:"task" binding:"required"`
	Status bool `json:"status" binding:"required"`
	UserId string `json:"userId" binding:"required"`
	Priority priority `json:"priority" binding:"required"`
}

// List of all Todos
var todos = []todo{}

// return All Todos
func getAllTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todos)
}

// return All Todos of Specific UserId
func getAllTodosOfUser(c *gin.Context) {
	userId := c.Param("userId")
	var userTodos = []todo{}

	for _, t := range todos {
		if t.UserId == userId {
			userTodos = append(userTodos, t)
		}
	}
	c.IndentedJSON(http.StatusOK, userTodos)
}

// return specific todo by id
func getTodoById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"id should be integer"})
		return
	}
	for _, t := range todos {
		if t.Id == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
}

// Delete Todo
func deleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"));
	userId := c.Param("userId")

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"id should be integer"})
		return
	}

	for _, t := range todos {
		if t.Id == id {
			if t.UserId == userId {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				return
			}
			var newTodos = deleteFilter(todos, func(i int) bool {return i != id})
			todos = newTodos
			c.IndentedJSON(http.StatusOK, todos);
			return
		}
	}
	c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Id not found"})
}

// return list removing specific todo
func deleteFilter(s []todo, fn func(int) bool) []todo {
	var t []todo
	for _, v := range s {
		if fn(v.Id) {
			t = append(t, v)
		}
	}
	return t
}


func main() {
	router := gin.Default()
	router.GET("/", getAllTodos)
	router.GET("/:userId", getAllTodosOfUser)
	router.GET("/todo/:id", getTodoById)
	router.DELETE("/todo/:id/:userId", deleteTodo)
	router.Run("localhost:3000")
	
}