import (
    "encoding/json"
    "net/http"
    "fmt"
)

func fetchCommits(owner, repo string) ([]map[string]interface{}, error) {
    url := fmt.Sprintf("https://api.github.com/repos/%s/%s/commits", owner, repo)

    resp, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    var data []map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
        return nil, err
    }

    return data, nil
}
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