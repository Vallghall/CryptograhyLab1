package main

import (
	caesar "firstlab/pkg/substitution"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Decipher result \t\t\t\t\t\t\t Cipher result")
	for key := 1; key < 26; key++ {
		fmt.Printf("%v\t\t%v\n", caesar.Decipher(os.Args[1], key), caesar.Cipher(os.Args[1], key))
	}
}
