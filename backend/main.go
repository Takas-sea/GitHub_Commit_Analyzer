package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gin-gonic/gin"
)

type Commit struct {
    Commit struct {
        Author struct {
            Date string `json:"date"`
        } `json:"author"`
        Message string `json:"message"`
    } `json:"commit"`
}

func fetchCommits(owner, repo string) ([]Commit, error) {
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", owner, repo)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("GitHub API error: %s", resp.Status)
    }

    var data []Commit
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }

    return data, nil
}

func main() {
    r := gin.Default()

    r.GET("/users/:username/stats", func(c *gin.Context) {
        username := c.Param("username")

        commits, err := fetchCommits(username, "your-repo-name")
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }

        c.JSON(200, gin.H{
            "raw": commits,
        })
    })

    r.Run(":8080")
}