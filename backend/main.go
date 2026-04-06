package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "pong",
        })
    })

    r.GET("/users/:username/stats", func(c *gin.Context) {
        username := c.Param("username") 
        c.JSON(http.StatusOK, gin.H{
            "username": username,
            "message":  "ここにGitHub APIから取得したデータを入れる予定",
        })
    })

    r.Run(":8080")
}