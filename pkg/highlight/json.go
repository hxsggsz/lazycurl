package highlight

import (
	"regexp"
	"strings"
)

const (
	COLOR_KEY       = "\033[34;1m" // Azul
	COLOR_STRING    = "\033[32m"   // Verde
	COLOR_NUMBER    = "\033[33m"   // Amarelo
	COLOR_BOOL_NULL = "\033[35m"   // Magenta
	COLOR_BRACKET   = "\033[37m"   // Branco
	COLOR_RESET     = "\033[0m"
)

func Json(input string) string {
	lines := strings.Split(input, "\n")

	for i, line := range lines {
		if !strings.Contains(line, ": ") {
			lines[i] = colorizeBrackets(line)
			continue
		}

		parts := strings.SplitN(line, ": ", 2)
		keyPart := parts[0]
		valPart := parts[1]

		keyPart = regexp.MustCompile(`"(.+)"`).ReplaceAllString(keyPart, COLOR_KEY+"\"$1\""+COLOR_RESET)

		valPart = colorizeValue(valPart)

		lines[i] = keyPart + ": " + valPart
	}

	return strings.Join(lines, "\n")
}

func colorizeValue(val string) string {
	if strings.Contains(val, "\"") {
		return regexp.MustCompile(`"([^"]*)"`).ReplaceAllString(val, COLOR_STRING+"\"$1\""+COLOR_RESET)
	}

	rePrimitivos := regexp.MustCompile(`\b(true|false|null|[0-9\.]+)\b`)

	if regexp.MustCompile(`\b(true|false|null)\b`).MatchString(val) {
		return rePrimitivos.ReplaceAllString(val, COLOR_BOOL_NULL+"$1"+COLOR_RESET)
	}

	return rePrimitivos.ReplaceAllString(val, COLOR_NUMBER+"$1"+COLOR_RESET)
}

func colorizeBrackets(line string) string {
	re := regexp.MustCompile(`([\{\}\[\]])`)
	return re.ReplaceAllString(line, COLOR_BRACKET+"$1"+COLOR_RESET)
}
