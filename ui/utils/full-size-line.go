package utils

import (
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/awesome-gocui/gocui"
)

var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

func FormatLineFullWidth(v *gocui.View, line string) string {
	width, _ := v.Size()

	// 1. Remove as cores apenas para calcular o tamanho visual
	plainText := ansiRegex.ReplaceAllString(line, "")

	// 2. CORREÇÃO: Conta os caracteres reais (Runes) e não os bytes
	visibleLen := utf8.RuneCountInString(plainText)

	// 3. Calcula o padding necessário
	paddingCount := width - visibleLen - 1
	if paddingCount <= 0 {
		return line
	}

	return line + strings.Repeat(" ", paddingCount)
}
