package util

import (
	"log"
	"os"
)

// EnsureFolder ensures that the folder path exists, creates it if not
func EnsureFolder(folder string) {
	if err := os.MkdirAll(folder, 0755); err != nil {
		log.Fatalf("failed to create folder %q: %v", folder, err)
	}
}
