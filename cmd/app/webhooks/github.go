package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/MartinHeinz/go-github-app/cmd/app/config"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func VerifySignature(payload []byte, signature string) bool {
	key := hmac.New(sha256.New, []byte(config.Config.GitHubWebhookSecret))
	key.Write([]byte(string(payload)))
	computedSignature := "sha256=" + hex.EncodeToString(key.Sum(nil))

	if computedSignature != signature {
		return false
	}
	return true
}

func ConsumeEvent(c *gin.Context) {
	payload, _ := ioutil.ReadAll(c.Request.Body)

	if !VerifySignature(payload, c.GetHeader("X-Hub-Signature-256")) {
		c.AbortWithStatus(http.StatusUnauthorized)
		log.Println("signatures don't match")
	}

	event := c.GetHeader("X-GitHub-Event")

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
