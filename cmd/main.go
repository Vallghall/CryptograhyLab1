package main

import (
	"firstlab/pkg/replacement"
	subst "firstlab/pkg/substitution"
	"fmt"
)

func main() {
	configs := NewConfigs()
	switch configs.method {
	case 1:
		if configs.decipher {
			fmt.Println(replacement.Decipher(configs.txt, configs.key))
		} else {
			fmt.Println(replacement.Cipher(configs.txt, configs.key))
		}
	case 2:
		if configs.decipher {
			fmt.Println(subst.Decipher(configs.txt, configs.key))
		} else {
			fmt.Println(subst.Cipher(configs.txt, configs.key))
		}
	}
}
