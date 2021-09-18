package substitution

import (
	"strings"
)

func Cipher(text string, key int) string {
	byteText := []byte(text)
	sb := strings.Builder{}
	for _, symbol := range byteText {
		symbol += byte(key)
		if symbol > byte('z') {
			sb.WriteByte(symbol - 26)
		} else {
			sb.WriteByte(symbol)
		}
	}
	return sb.String()
}

func Decipher(text string, key int) string {
	byteText := []byte(text)
	sb := strings.Builder{}
	for _, symbol := range byteText {
		symbol -= byte(key)
		if symbol < byte('a') {
			sb.WriteByte(symbol + 26)
		} else {
			sb.WriteByte(symbol)
		}
	}
	return sb.String()
}
