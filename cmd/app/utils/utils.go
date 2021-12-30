package utils

import (
	"github.com/MartinHeinz/go-github-app/cmd/app/config"
	"github.com/bradleyfalzon/ghinstallation/v2"
	"github.com/google/go-github/v41/github"
	"log"
	"net/http"
)

func InitGitHubClient() {
	tr := http.DefaultTransport
	itr, err := ghinstallation.NewKeyFromFile(tr, APP_ID, INSTALLATION_ID, "/config/github-app.pem")

	if err != nil {
		log.Fatal(err)
	}

	config.Config.GitHubClient = github.NewClient(&http.Client{Transport: itr})
}
