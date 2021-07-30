package main 

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Welcome to the world of Go and Gin")
	})
	r.Run(":8090")
}
