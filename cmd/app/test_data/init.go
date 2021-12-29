package test_data

import (
	"github.com/MartinHeinz/go-github-app/cmd/app/config"
	"github.com/MartinHeinz/go-github-app/cmd/app/utils"
)

// Initializes application config and SQLite database used for testing
func init() {
	// the test may be started from the home directory or a subdirectory
	err := config.LoadConfig("/config") // on host use absolute path
	if err != nil {
		panic(err)
	}

	utils.InitGitHubClient()
}

func GetTestCaseFolder() string {
	return "/test_data/test_case_data" // on host use absolute path
}
