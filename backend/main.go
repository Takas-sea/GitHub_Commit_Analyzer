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

func calculateScore(daily map[string]int) (int, int, int) {
    total := 0
    longest := 0
    current := 0

    for _, count := range daily {
        total += count
        if count > 0 {
            current++
            if current > longest {
                longest = current
            }
        } else {
            current = 0
        }
    }

    score := (longest * 2) + total
    return total, longest, score
}

func aggregateByDate(commits []Commit) map[string]int {
    result := make(map[string]int)

    for _, c := range commits {
        date := c.Commit.Author.Date[:10] // YYYY-MM-DD
        result[date]++
    }

    return result
}

func classify(msg string) string {
    if strings.HasPrefix(msg, "feat") {
        return "feature"
    }
    if strings.HasPrefix(msg, "fix") {
        return "bugfix"
    }
    return "other"
}

func main() {
    r := gin.Default()

    r.GET("/users/:username/stats", func(c *gin.Context) {
        username := c.Param("username")
        repo := c.Query("repo") // ?repo=xxx

if repo == "" {
    c.JSON(400, gin.H{"error": "repo is required"})
    return
}
        commits, err := fetchCommits(username, repo)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }

       daily := aggregateByDate(commits)

c.JSON(200, gin.H{
    "raw":   commits,
    "daily": daily,
})
    })

    r.Run(":8080")
}