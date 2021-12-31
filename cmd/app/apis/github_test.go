package apis

import (
	"github.com/MartinHeinz/go-github-app/cmd/app/config"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v41/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func TestGithubGetPullRequests(t *testing.T) {
	expectedTitles := []string{
		"PR number one",
		"PR number three",
	}
	stateOpen := "open"
	stateClosed := "closed"
	closedPullRequestTitle := "PR number two"
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetReposPullsByOwnerByRepo,
			[]github.PullRequest{
				github.PullRequest{
					State: &stateOpen,
					Title: &expectedTitles[0],
				},
				github.PullRequest{
					State: &stateClosed,
					Title: &closedPullRequestTitle,
				},
				github.PullRequest{
					State: &stateOpen,
					Title: &expectedTitles[1],
				},
			},
		),
	)
	client := github.NewClient(mockedHTTPClient)
	config.Config.GitHubClient = client

	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(res)
	ctx.Params = []gin.Param{
		gin.Param{Key: "owner", Value: "octocat"},
		gin.Param{Key: "repo", Value: "hello-world"},
	}

	GetPullRequests(ctx)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		println(err)
	}

	assert.Equal(t, 200, res.Code)
	assert.Contains(t, string(body), expectedTitles[0])
	assert.NotContains(t, string(body), closedPullRequestTitle[1])
	assert.Contains(t, string(body), expectedTitles[1])
}
