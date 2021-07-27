package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tsvillain/go-todo-server/db"
	"github.com/tsvillain/go-todo-server/entity"
)

// return All Todos
func GetAllTodos(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, db.Todos)
}

// return All Todos of Specific username
func GetAllTodosOfUser(c *gin.Context) {
	username := c.Param("username")
	userTodos := allTodoOfSpecificUser(username)
	c.IndentedJSON(http.StatusOK, userTodos)
}

// return specific todo by id
func GetTodoById(c *gin.Context) {
	id, _ := getIdFromParam(c)
	for _, t := range db.Todos {
		if t.Id == id {
			c.IndentedJSON(http.StatusOK, t)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Todo found for ID = " + fmt.Sprint(id)})
}

// Add todo
func AddTodo(c *gin.Context) {
	var newTodo entity.Todo
	err := c.BindJSON(&newTodo)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	validateRequest(c, newTodo.UserName, newTodo.Task)
	db.Todos = append(db.Todos, newTodo)
	userTodos := allTodoOfSpecificUser(newTodo.UserName)
	c.IndentedJSON(http.StatusCreated, userTodos)
}

// Delete Todo
func DeleteTodo(c *gin.Context) {
	id, _ := getIdFromParam(c)
	username := c.Param("username")
	notFound := true

	for _, t := range db.Todos {
		if t.Id == id {
			notFound = false
			if t.UserName != username {
				c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
				c.Abort()
			}
			newTodos := deleteFilter(db.Todos, func(i int, j string) bool { return i != id })
			db.Todos = newTodos
			userTodos := allTodoOfSpecificUser(t.UserName)
			c.IndentedJSON(http.StatusOK, userTodos)
			c.Abort()
		}
	}
	if notFound {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "No Todo Found"})
	}
}

// update existing todo throws error if not existing todo found
func UpdateTodo(c *gin.Context) {
	var updatedTodo entity.Todo
	err := c.BindJSON(&updatedTodo)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	validateRequest(c, updatedTodo.UserName, updatedTodo.Task)
	todoAfterDelete := deleteFilter(db.Todos, func(i int, j string) bool { return !(i == updatedTodo.Id && j == updatedTodo.UserName) })
	if len(todoAfterDelete) == len(db.Todos) {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No Todo Exist, Can't Update"})
		c.Abort()
	}
	todoAfterDelete = append(todoAfterDelete, updatedTodo)
	db.Todos = todoAfterDelete
	userTodos := allTodoOfSpecificUser(updatedTodo.UserName)
	c.IndentedJSON(http.StatusOK, userTodos)

}

// return list removing specific todo
func deleteFilter(s []entity.Todo, fn func(int, string) bool) []entity.Todo {
	var t []entity.Todo
	for _, v := range s {
		if fn(v.Id, v.UserName) {
			t = append(t, v)
		}
	}
	return t
}

func validateRequest(c *gin.Context, username string, task string) {
	if username == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Username is requried"})
		c.Abort()
	}
	if task == "" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task is requried"})
		c.Abort()
	}

}

// get All todo of specific user
func allTodoOfSpecificUser(username string) []entity.Todo {
	var userTodos = []entity.Todo{}
	for _, t := range db.Todos {
		if t.UserName == username {
			userTodos = append(userTodos, t)
		}
	}
	return userTodos
}

/// utils
// get id from param
func getIdFromParam(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return 0, err
	}

	return id, nil
}
