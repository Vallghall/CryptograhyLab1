package substitution

import (
	"strings"
	"unicode"
)

func Cipher(text string, key int) string {
	sb := strings.Builder{}
	for _, symbol := range text {
		if unicode.IsLetter(symbol) {
			symbol += rune(key)
			if unicode.ToUpper(symbol) > 'Z' {
				sb.WriteRune(symbol - rune(26))
			} else {
				sb.WriteRune(symbol)
			}
		} else {
			sb.WriteRune(symbol)
		}
	}
	return sb.String()
}

func Decipher(text string, key int) string {
	sb := strings.Builder{}
	for _, symbol := range text {
		if unicode.IsLetter(symbol) {
			symbol -= rune(key)
			if unicode.ToUpper(symbol) < 'A' {
				sb.WriteRune(symbol + rune(26))
			} else {
				sb.WriteRune(symbol)
			}
		} else {
			sb.WriteRune(symbol)
		}
	}
	return sb.String()
}
