package test_data

import (
	"github.com/MartinHeinz/go-github-app/cmd/app/config"
)

// Initializes application config
func init() {
	// the test may be started from the home directory or a subdirectory
	err := config.LoadConfig("/config") // on host use absolute path
	if err != nil {
		panic(err)
	}
}

func GetTestCaseFolder() string {
	return "/test_data/test_case_data" // on host use absolute path
}
