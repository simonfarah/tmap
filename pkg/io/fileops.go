package io

import (
	"encoding/json"
	"os"
)

func ReadJSONFile(filePath string) (map[string]interface{}, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func WriteJSONFile(filePath string, content interface{}) error {
	data, err := json.Marshal(content)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}
