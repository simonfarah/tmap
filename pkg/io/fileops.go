package io

import (
	"fmt"
	"os"
)

func ReadFile(filePath string) []byte {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	return content
}

func WriteFile(filePath string, content interface{}) {
	data := []byte(fmt.Sprintf("%v", content))

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}
}
