package process

import "strings"

func EncodeGroup(group []rune, key int) string {
	sb := strings.Builder{}
	ks := group[0] + rune(key)
	sb.WriteRune(ks)
	for _, symbol := range group[1:] {
		temp := (symbol + ks - 'a' - 'a') % 26
		sb.WriteRune(temp + 'a')
	}
	return sb.String()
}

func DecodeGroup(group []rune, key int) string {
	sb := strings.Builder{}
	ks := group[0]
	sb.WriteRune(ks - rune(key))
	for _, symbol := range group[1:] {
		temp := symbol - ks
		if temp < 0 {
			temp += 26
		}
		sb.WriteRune(temp + 'a')
	}
	return sb.String()
}
