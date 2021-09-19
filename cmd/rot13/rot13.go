package main

import (
	caesar "firstlab/pkg/substitution"
	"fmt"
	"os"
)

func main() {
	fmt.Println(caesar.Cipher(os.Args[1], 13))
}
