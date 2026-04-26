package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/users/:username/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"total_commits":  10,
			"longest_streak": 3,
			"score":          16,
			"daily_commits":  []interface{}{},
		})
	})

	r.Run(":8080")
}