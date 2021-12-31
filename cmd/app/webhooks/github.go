package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
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

var Consumers = map[string]func(EventPayload) error{
	string(Install): consumeInstallEvent,
	Ping:            consumePingEvent,
	Push:            consumePushEvent,
	PullRequest:     consumePullRequestEvent,
}

func VerifySignature(payload []byte, signature string) bool {
	key := hmac.New(sha256.New, []byte(config.Config.GitHubWebhookSecret))
	key.Write([]byte(string(payload)))
	computedSignature := "sha256=" + hex.EncodeToString(key.Sum(nil))
	log.Printf("computed signature: %s", computedSignature)

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
			found = true
			log.Printf("consuming event: %s", e)
			var p EventPayload
			json.Unmarshal(payload, &p)
			if err := Consumers[string(e)](p); err != nil {
				log.Printf("couldn't consume event %s, error: %+v", string(e), err)
				// We're responding to GitHub API, we really just want to say "OK" or "not OK"
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": err})
			}
			log.Printf("consumed event: %s", e)
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
	}
	if !found {
		log.Printf("Unsupported event: %s", event)
		c.AbortWithStatus(http.StatusNotImplemented)
		return
	}
	// Otherwise...
	// Add some error into response payload
	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"reason": "some error..."})
}

func consumeInstallEvent(payload EventPayload) error {
	// Process event ...
	// Insert data into database ...
	return nil
}

func consumePingEvent(payload EventPayload) error {
	// Process event ...
	// Insert data into database ...
	return nil
}

func consumePushEvent(payload EventPayload) error {

	// Process event ...
	// Insert data into database ...

	log.Printf("Received push from %s, by user %s, on branch %s",
		payload.Repository.FullName,
		payload.Pusher.Name,
		payload.Ref)

	// Enumerating commits
	var commits []string
	for _, commit := range payload.Commits {
		commits = append(commits, commit.ID)
	}
	log.Printf("Pushed commits: %v", commits)

	return nil
}

func consumePullRequestEvent(payload EventPayload) error {
	// Process event ...
	// Insert data into database ...
	return nil
}
