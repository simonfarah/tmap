package converters

import (
	"fmt"
	"strings"
)

func ConvertToStarshipPalette(colors map[string]interface{}, name string) string {
	formattedName := strings.NewReplacer(" ", "_", ".", "").Replace(strings.ToLower(name))

	return fmt.Sprintf(`
palette = "%v"

[palettes.%v]
black = "%v"
red = "%v"
green = "%v"
yellow = "%v"
blue = "%v"
purple = "%v"
cyan = "%v"
white = "%v"
# bright colors
bright_black = "%v"
bright_red = "%v"
bright_green = "%v"
bright_yellow = "%v"
bright_blue = "%v"
bright_purple = "%v"
bright_cyan = "%v"
bright_white = "%v"
`,
		formattedName,
		formattedName,
		colors["terminal.ansiBlack"],
		colors["terminal.ansiRed"],
		colors["terminal.ansiGreen"],
		colors["terminal.ansiYellow"],
		colors["terminal.ansiBlue"],
		colors["terminal.ansiMagenta"],
		colors["terminal.ansiCyan"],
		colors["terminal.ansiWhite"],
		colors["terminal.ansiBrightBlack"],
		colors["terminal.ansiBrightRed"],
		colors["terminal.ansiBrightGreen"],
		colors["terminal.ansiBrightYellow"],
		colors["terminal.ansiBrightBlue"],
		colors["terminal.ansiBrightMagenta"],
		colors["terminal.ansiBrightCyan"],
		colors["terminal.ansiBrightWhite"],
	)
}
