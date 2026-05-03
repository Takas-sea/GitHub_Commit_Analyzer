package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Repo struct {
	Name string `json:"name"`
}

func fetchRepos(username string) ([]Repo, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos?per_page=100", username)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var repos []Repo
	if err := json.NewDecoder(resp.Body).Decode(&repos); err != nil {
		return nil, err
	}

	return repos, nil
}

func main() {
	r := gin.Default()

	r.GET("/users/:username/stats", func(c *gin.Context) {
		username := c.Param("username")

		repos, err := fetchRepos(username)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to fetch repos"})
			return
		}

		c.JSON(200, gin.H{
			"repo_count": len(repos),
			"repos":      repos,
		})
	})

	r.Run(":8080")
}