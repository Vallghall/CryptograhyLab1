package polygrammar

import (
	"firstlab/internal/process"
	"strings"
)

func Cipher(text string, key int) string {
	sb := strings.Builder{}
	symbols := []rune(text)
	for i, j := 0, key; j <= len(text); i, j = i+key, j+key {
		sb.WriteString(process.EncodeGroup(symbols[i:j], key))
	}
	return sb.String()
}

func Decipher(text string, key int) string {
	sb := strings.Builder{}
	symbols := []rune(text)
	for i, j := 0, key; j <= len(text); i, j = i+key, j+key {
		sb.WriteString(process.DecodeGroup(symbols[i:j], key))
	}
	return sb.String()
}
