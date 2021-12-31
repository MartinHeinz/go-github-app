package webhooks

import (
	"github.com/MartinHeinz/go-github-app/cmd/app/test_data"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGithubVerifySignature(t *testing.T) {
	payload, _ := ioutil.ReadFile("/github_installation_input_t1.json")
	signature := "sha256=d77718c0b2f2a9ae94d84846ec2034914042facb274a3b0a23d4140cb2f1c0df"

	result := VerifySignature(payload, signature)
	assert.True(t, result)
}

func TestGithub(t *testing.T) {
	path := test_data.GetTestCaseFolder()
	runWebhookTests(t, []webhookTestCase{
		{
			"T1 - Consume Push Event",
			"POST",
			"/github/payload",
			"/github/payload",
			map[string]string{
				"X-GitHub-Event":      "push",
				"X-Hub-Signature-256": "sha256=412e84d3fc669078574abc0f9a7dcb1e3806182d4c333bdc8e09404523eb4a28",
			},
			path + "/github_push_input_t1.json",
			ConsumeEvent,
			http.StatusNoContent,
			""},
	})
}
