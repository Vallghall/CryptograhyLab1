package main

import (
	replace "firstlab/pkg/replacement"
	"flag"
	"fmt"
)

func main() {
	key, inputText := parseFlags()
	fmt.Println(replace.Cipher(inputText, key))
}

func parseFlags() (int, string) {
	key := flag.Int("key", 5, "Sets the key for the substitution/replacement cipher")
	txt := flag.String("text", "exampletxt", "Text for encoding")
	flag.Parse()
	return *key, *txt
}
