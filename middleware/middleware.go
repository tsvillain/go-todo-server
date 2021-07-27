package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tsvillain/go-todo-server/db"
)

// Check if Id is int else response with error
func ValidateId() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := getIdFromParam(c)
		if err != nil {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "id should be integer"})
			c.Abort()
		}
	}
}

// check if user exist
// current we don't have any user collection so if any todo exist with username then user exist.
func ValidateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		exist := false
		for _, t := range db.Todos {
			if t.UserName == username {
				exist = true
				break
			}
		}
		if !exist {
			c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "No user Exist with username = " + username})
			c.Abort()
		}
	}
}

/// utils

// get id from param
func getIdFromParam(c *gin.Context) error {
	_, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return err
	}

	return nil
}
