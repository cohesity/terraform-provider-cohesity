package utils

import "strings"

func EscapeSpecialSymbols(input string) string {
	specialSymbols := `\.$^*+?()[]{}|`

	var escaped strings.Builder
	for _, ch := range input {
		if strings.ContainsRune(specialSymbols, ch) {
			escaped.WriteRune('\\')
		}
		escaped.WriteRune(ch)
	}

	return escaped.String()
}
