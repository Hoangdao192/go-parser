package util

import (
	"io"
	"os"
)

func ReadFile(filePath string) (string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Read the contents of the file
	content, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
