package util

import (
	"os"
)

// Check if directory or file exists
func dirExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
