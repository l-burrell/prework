package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var nextId int = 0
var todos []Todo

func GetNextId() int {
	value := nextId
	nextId++
	return value
}

func GetTodos(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"list": todos})
}

func main() {
	todos = append(todos, Todo{Id: GetNextId(), Value: "CodeHouse", DueDate: "7/31/2021"})
	r := gin.Default()
	r.GET("/api/todos", GetTodos)
	r.Run(":8090")
}
