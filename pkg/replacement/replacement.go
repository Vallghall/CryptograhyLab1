package replacement

import "fmt"

func Cipher(text string, key int) string {
	matrix := newMatrix(len(text), key)
	for _, symbol := range text {
		matrix.addToMatrix(symbol)
	}
	return fmt.Sprintln(matrix.transformMatrix())
}

/*
func Decipher(text string, key int) string {
	matrix := newMatrix(len(text))
	for _, symbol := range text {
		matrix.addToMatrix(symbol, key)
	}
	return fmt.Sprintln(matrix.transformMatrix(key))
}
*/
