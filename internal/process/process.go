package process

import "strings"

func EncodeGroup(group []rune, key int) string {
	sb := strings.Builder{}
	ks := group[0] + rune(key)
	sb.WriteRune(ks)
	for _, symbol := range group[1:] {
		sb.WriteRune(symbol + ks - 'a')
	}
	return sb.String()
}

func DecodeGroup(group []rune, key int) string {
	sb := strings.Builder{}
	ks := group[0]
	sb.WriteRune(ks - rune(key))
	for _, symbol := range group[1:] {
		sb.WriteRune(symbol - ks + 'a')
	}
	return sb.String()
}
