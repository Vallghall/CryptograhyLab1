package replacement

import (
	table "firstlab/internal/matrix"
	"fmt"
)

func Cipher(text string, key int) string {
	matrix := table.NewMatrix(len(text), key)
	for _, symbol := range text {
		matrix.AddToMatrix(symbol)
	}
	return fmt.Sprintln(matrix.TransformMatrix())
}
