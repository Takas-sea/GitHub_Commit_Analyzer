package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users/:username/stats", func(c *gin.Context) {
    username := c.Param("username")

    c.JSON(http.StatusOK, gin.H{
    "username": username,
	})
	})
	r.Run(":8080")
}