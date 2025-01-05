package themes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

type Theme struct {
	Name        string
	Path        string
	IsDarkTheme bool
}

func ExtractVsThemes() ([]Theme, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to find home directory: %w", err)
	}

	extensionsDir := filepath.Join(homeDir, ".vscode", "extensions")
	if _, err := os.Stat(extensionsDir); os.IsNotExist(err) {
		return nil, errors.New("extensions directory not found")
	}

	var themes []Theme
	dirs, err := os.ReadDir(extensionsDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read extensions directory: %w", err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		packageJSONPath := filepath.Join(extensionsDir, dir.Name(), "package.json")
		if _, err := os.Stat(packageJSONPath); os.IsNotExist(err) {
			continue
		}

		data, err := os.ReadFile(packageJSONPath)
		if err != nil {
			fmt.Printf("Failed to read package.json for %s: %v\n", dir.Name(), err)
			continue
		}

		var packageJSON struct {
			Contributes struct {
				Themes []struct {
					Label   string `json:"label"`
					UITheme string `json:"uiTheme"`
					Path    string `json:"path"`
				} `json:"themes"`
			} `json:"contributes"`
		}

		if err := json.Unmarshal(data, &packageJSON); err != nil {
			fmt.Printf("Failed to parse package.json for %s: %v\n", dir.Name(), err)
			continue
		}

		for _, theme := range packageJSON.Contributes.Themes {
			themes = append(themes, Theme{
				Name:        theme.Label,
				Path:        filepath.Join(extensionsDir, dir.Name(), theme.Path),
				IsDarkTheme: theme.UITheme == "vs-dark" || theme.UITheme == "hc-black",
			})
		}
	}

	return themes, nil
}
