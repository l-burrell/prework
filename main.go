package main 

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
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

func PostTodo(c *gin.Context) {
	var item Todo
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item.Id = GetNextId()
	todos = append(todos, item)
	c.String(http.StatusCreated, c.FullPath() + "/" + strconv.Itoa(item.Id))
}

func main() {
	todos = append(todos, Todo{Id: GetNextId(), Value: "CodeHouse", DueDate: "7/31/2021"})
	r := gin.Default()
	r.GET("/api/todos", GetTodos)
	r.POST("/api/todos", PostTodo)
	r.Run(":8090")
}
