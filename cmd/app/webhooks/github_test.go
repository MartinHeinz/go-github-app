package webhooks

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestGithubVerifySignature(t *testing.T) {
	payload, _ := ioutil.ReadFile("/github_installation_input_t1.json")
	signature := "sha256=b613679a0814d9ec772f95d778c35fc5ff1697c493715653c6c712144292c5ad"

	result := VerifySignature(payload, signature)
	assert.True(t, result)
}
