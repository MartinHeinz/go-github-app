package test_data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
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
