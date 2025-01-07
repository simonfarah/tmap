package io

import (
	"encoding/json"
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
	var data []byte
	var err error

	switch v := content.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	case map[string]interface{}:
		data, err = json.Marshal(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error marshalling content: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unsupported content type: %T\n", v)
		os.Exit(1)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
		os.Exit(1)
	}
}
