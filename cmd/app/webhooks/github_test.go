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
	signature := "sha256=28377c28d5e1cabe15e0743189855e0fd04e4ad4643083d48608bd0e0ddfaa1d"

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
				"X-Hub-Signature-256": "sha256=f86efe4b911eeb240996958e9877fe294c5057cb2a433b5f8d161678e19a52af",
			},
			path + "/github_push_input_t1.json",
			ConsumeEvent,
			http.StatusNoContent,
			""},
	})
}
