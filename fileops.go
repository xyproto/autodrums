package main

import (
	"os"
	"path/filepath"
	"strings"
)

// findFiles searches for files with a specific extension in a given directory.
func findFiles(rootPath, ext string) ([]string, error) {
	var files []string
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

// iContains checks if a string contains the specified substring, case-insensitive.
func iContains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
