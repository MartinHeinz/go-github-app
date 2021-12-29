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

type Event string

const (
	Install     Event = "installation"
	Ping              = "ping"
	Push              = "push"
	PullRequest       = "pull_request"
)

var Events = []Event{
	Install,
	Ping,
	Push,
	PullRequest,
}

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

	found := false
	for _, e := range Events {
		if string(e) == event {
			log.Printf("Consume %s", e)
			found = true
		}
	}
	if !found {
		log.Printf("Unsupported event: %s", event)
	}

	c.JSON(http.StatusOK, gin.H{
		"event": event,
	})
}
