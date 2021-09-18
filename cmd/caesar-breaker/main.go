package main

import (
	caesar "firstlab/pkg/substitution"
	"fmt"
	"os"
)

func main() {
	for key := 1; key < 26; key++ {
		fmt.Println(caesar.Decipher(os.Args[1], key))
	}
}
