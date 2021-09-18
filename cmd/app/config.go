package main

import "flag"

type Configs struct {
	key      int
	txt      string
	method   int
	decipher bool
}

func NewConfigs() *Configs {
	key := flag.Int("key", 5, "Sets the key for the substitution/replacement cipher")
	txt := flag.String("text", "exampletxt", "Text for encoding")
	method := flag.Int("c", 1, "Choice of cryptography method")
	decipher := flag.Bool("d", false, "Decrypts input when used")
	flag.Parse()
	return &Configs{
		key:      *key,
		txt:      *txt,
		method:   *method,
		decipher: *decipher,
	}
}
