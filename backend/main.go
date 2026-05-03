package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Repo struct {
	Name string `json:"name"`
}

type Commit struct {
	Commit struct {
		Author struct {
			Date string `json:"date"`
		} `json:"author"`
		Message string `json:"message"`
	} `json:"commit"`
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

func fetchCommits(owner, repo string) ([]Commit, error) {
	since := time.Now().AddDate(0, 0, -30).UTC().Format(time.RFC3339)

	url := fmt.Sprintf(
		"https://api.github.com/repos/%s/%s/commits?since=%s",
		owner,
		repo,
		since,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var commits []Commit
	if err := json.NewDecoder(resp.Body).Decode(&commits); err != nil {
		return nil, err
	}

	return commits, nil
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

		if len(repos) == 0 {
			c.JSON(200, gin.H{"message": "no repos"})
			return
		}

		commits, err := fetchCommits(username, repos[0].Name)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to fetch commits"})
			return
		}

		c.JSON(200, gin.H{
			"repo":    repos[0].Name,
			"commits": commits,
		})
	})

	r.Run(":8080")
}