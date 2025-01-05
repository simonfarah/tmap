package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/simonfarah/tmap/internal/converters"
	"github.com/simonfarah/tmap/internal/themes"
	"github.com/simonfarah/tmap/pkg/io"

	"github.com/AlecAivazis/survey/v2"
)

func main() {
	installedThemes, err := themes.ExtractVsThemes()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting themes: %v\n", err)
		os.Exit(1)
	}

	if len(installedThemes) == 0 {
		fmt.Println("No VSCode themes found")
		os.Exit(1)
	}

	choices := make([]string, len(installedThemes))
	for i, theme := range installedThemes {
		choices[i] = theme.Name
	}

	var selectedTheme string
	prompt := &survey.Select{
		Message:  "Choose a theme:",
		Options:  choices,
		PageSize: 5,
	}

	if err := survey.AskOne(prompt, &selectedTheme); err != nil {
		fmt.Fprintf(os.Stderr, "Prompt failed: %v\n", err)
		os.Exit(1)
	}

	var themeDetails *themes.Theme
	for _, theme := range installedThemes {
		if theme.Name == selectedTheme {
			themeDetails = &theme
			break
		}
	}

	if themeDetails == nil {
		fmt.Fprintf(os.Stderr, "Selected theme not found\n")
		os.Exit(1)
	}

	themeFileData, err := io.ReadJSONFile(themeDetails.Path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading theme file: %v\n", err)
		os.Exit(1)
	}

	themeColors, ok := themeFileData["colors"].(map[string]interface{})
	if !ok {
		fmt.Fprintf(os.Stderr, "Error parsing theme colors\n")
		os.Exit(1)
	}

	windowsTerminalScheme := converters.ConvertToWindowsTerminalScheme(themeColors, themeDetails.Name)
	windowsTerminalTheme := converters.ConvertToWindowsTerminalTheme(themeColors, themeDetails.Name, themeDetails.IsDarkTheme)

	if err := os.MkdirAll("generated", os.ModePerm); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating 'generated' directory: %v\n", err)
		os.Exit(1)
	}

	themeDir := filepath.Join("generated", themeDetails.Name)
	if err := os.MkdirAll(themeDir, os.ModePerm); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating theme directory: %v\n", err)
		os.Exit(1)
	}

	windowsTerminalSchemeFile := filepath.Join(themeDir, "windows-terminal-scheme.json")
	windowsTerminalThemeFile := filepath.Join(themeDir, "windows-terminal-theme.json")

	if err := io.WriteJSONFile(windowsTerminalSchemeFile, windowsTerminalScheme); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing scheme file: %v\n", err)
		os.Exit(1)
	}

	if err := io.WriteJSONFile(windowsTerminalThemeFile, windowsTerminalTheme); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing theme file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Files generated successfully")
}
