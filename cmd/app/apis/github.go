package apis

import (
	"github.com/MartinHeinz/go-github-app/cmd/app/config"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v41/github"
	"log"
	"net/http"
)

// TODO Pagination
// Manual Test: curl http://localhost:8080/api/v1/github/pullrequests/MartinHeinz/python-project-blueprint
// Result `{"pull_requests":["Some Instructions","Add newline to match dev.Dockerfile"]}`
func GetPullRequests(c *gin.Context) {
	owner := c.Param("owner")
	repo := c.Param("repo")
	if pullRequests, resp, err := config.Config.GitHubClient.PullRequests.List(c, owner, repo, &github.PullRequestListOptions{
		State: "open",
	}); err != nil {
		c.AbortWithStatus(resp.StatusCode)
		log.Println(err)
	} else {
		var pullRequestTitles []string
		for _, pr := range pullRequests {
			pullRequestTitles = append(pullRequestTitles, *pr.Title)
		}
		c.JSON(http.StatusOK, gin.H{
			"pull_requests": pullRequestTitles,
		})
	}
}
