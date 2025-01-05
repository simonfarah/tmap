package converters

func ConvertToWindowsTerminalScheme(colors map[string]interface{}, name string) map[string]interface{} {
	return map[string]interface{}{
		"name": name,

		"background": colors["editor.background"],
		"foreground": colors["editor.foreground"],

		"cursorColor": colors["editorCursor.foreground"],

		"black":  colors["terminal.ansiBlack"],
		"red":    colors["terminal.ansiRed"],
		"green":  colors["terminal.ansiGreen"],
		"yellow": colors["terminal.ansiYellow"],
		"blue":   colors["terminal.ansiBlue"],
		"purple": colors["terminal.ansiMagenta"],
		"cyan":   colors["terminal.ansiCyan"],
		"white":  colors["terminal.ansiWhite"],

		"brightBlack":  colors["terminal.ansiBrightBlack"],
		"brightRed":    colors["terminal.ansiBrightRed"],
		"brightGreen":  colors["terminal.ansiBrightGreen"],
		"brightYellow": colors["terminal.ansiBrightYellow"],
		"brightBlue":   colors["terminal.ansiBrightBlue"],
		"brightPurple": colors["terminal.ansiBrightMagenta"],
		"brightCyan":   colors["terminal.ansiBrightCyan"],
		"brightWhite":  colors["terminal.ansiBrightWhite"],
	}
}

func ConvertToWindowsTerminalTheme(colors map[string]interface{}, name string, isDarkTheme bool) map[string]interface{} {
	return map[string]interface{}{
		"name": name,

		"tab": map[string]interface{}{
			"background":          colors["tab.activeBackground"],
			"unfocusedBackground": nil,
			"showCloseButton":     "always",
		},

		"tabRow": map[string]interface{}{
			"background":          colors["titleBar.activeBackground"],
			"unfocusedBackground": colors["titleBar.inactiveBackground"],
		},

		"window": map[string]interface{}{
			"applicationTheme": func() string {
				if isDarkTheme {
					return "dark"
				}
				return "light"
			}(),
		},
	}
}
