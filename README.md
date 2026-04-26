# GitHub Commit Analyzer

## Goal

To understand the continuity of development and activity trends.

## Feature

* Retrieve commit history
* Commit Classification
* Graph display
* Scoring function

## Architecture

```mermaid
graph TD
    A["Frontend (Next.js)"] --> |HTTP Request| B["Backend (Go + Gin)"]
    B --> C["GitHub API"]
```

## MVP Specification

### Scope

* Target: GitHub user (public repositories only)
* Period: Last 30 days
* Authentication: Not required (public data only)

---

### Backend (Go)

#### Technology Stack
- Go 1.21+
- Web framework: Gin
- HTTP Client: `net/http` or `github.com/go-resty/resty/v2`
- JSON handling: `encoding/json`
- Concurrency: goroutines + channels (optional, for faster API calls)

#### Endpoint

GET /users/{username}/stats

#### Processing

1. Retrieve commit history from GitHub API
2. Aggregate commits by date (daily count)
3. Classify commit messages:

   * "feat" → feature
   * "fix" → bugfix
   * others → other
4. Calculate score:
   score = (longest streak × 2) + total commits


### Frontend (Next.js)

#### UI

* Input field for GitHub username
* Button to fetch data

#### Display

* Bar chart of daily commits (last 30 days)
* Simple display of:

  * total commits
  * score
---

### API

### GET /users/{username}/stats

#### Response
{
  "total_commits": number,
  "longest_streak": number,
  "score": number,
  "daily_commits": [
    { "date": "YYYY-MM-DD", "count": number }
  ]
}

---

### Data Fetch Strategy
- Fetch all public repositories of the user
- Retrieve commits per repository within the last 30 days
- Aggregate results on backend

---

### Out of Scope

* Authentication (OAuth)
* Database (PostgreSQL)
* Private repositories
* Advanced analytics
* Selecting specific repositories instead of analyzing all public repositories
* UI optimization

---

### Definition of Done

* User can input a GitHub username
* Commit data is fetched and displayed
* Daily commit graph is visible
* Score is calculated and shown
