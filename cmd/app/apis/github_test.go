package apis

import (
	"github.com/MartinHeinz/go-github-app/cmd/app/config"
	"github.com/gin-gonic/gin"
	"github.com/google/go-github/v41/github"
	"github.com/migueleliasweb/go-github-mock/src/mock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGithubGetPullRequests(t *testing.T) {
	expectedTitles := []string{
		"PR number one",
		"PR number three",
	}
	closedPullRequestTitle := "PR number two"
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatch(
			mock.GetReposPullsByOwnerByRepo,
			[]github.PullRequest{
				{
					State: github.String("open"),
					Title: &expectedTitles[0],
				},
				{
					State: github.String("closed"),
					Title: &closedPullRequestTitle,
				},
				{
					State: github.String("open"),
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
		{Key: "owner", Value: "octocat"},
		{Key: "repo", Value: "hello-world"},
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

func TestGithubGetPullRequests_Error(t *testing.T) {
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatchHandler(
			mock.GetReposPullsByOwnerByRepo,
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				mock.WriteError(
					w,
					http.StatusInternalServerError,
					"GitHub downtime...",
				)
			}),
		),
	)
	client := github.NewClient(mockedHTTPClient)
	config.Config.GitHubClient = client

	gin.SetMode(gin.TestMode)
	res := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(res)
	ctx.Params = []gin.Param{
		{Key: "owner", Value: "octocat"},
		{Key: "repo", Value: "hello-world"},
	}

	GetPullRequests(ctx)

	assert.Equal(t, http.StatusInternalServerError, res.Code)
}

func TestGithubGetPullRequestsPaginated(t *testing.T) {
	expectedTitles := []string{
		"PR number one - first page",
		"PR number two - first page",
		"PR number three - second page",
	}
	mockedHTTPClient := mock.NewMockedHTTPClient(
		mock.WithRequestMatchPages(
			// Mock completely overrides `PullRequestListOptions` so any additional args, such as `State` will disappear
			mock.GetReposPullsByOwnerByRepo,
			[]github.PullRequest{
				{
					Title: &expectedTitles[0],
				},
				{
					Title: &expectedTitles[1],
				},
			},
			[]github.PullRequest{
				{
					Title: &expectedTitles[2],
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
		{Key: "owner", Value: "octocat"},
		{Key: "repo", Value: "hello-world"},
		{Key: "page", Value: "2"},
	}

	GetPullRequestsPaginated(ctx)

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		println(err)
	}

	assert.Equal(t, 200, res.Code)
	assert.Contains(t, string(body), expectedTitles[0])
	assert.Contains(t, string(body), expectedTitles[1])
	assert.Contains(t, string(body), expectedTitles[2])
}
