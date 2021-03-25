// THIS IS FOR NOTE!!!:
// An Error which occurs when input contains spaces needs to be handled!!! (status: DONE!)
// Create Matrix structure (status: DONE)
// MatrixPrint function (status: DONE)
// ERROR splitting input, empties still remaining in some cases!!! (status: TODO)
// Check for compatibility of matrixes before operetaions! (status: TODO)
// Implement operations (status: TODO)
package main

import (
	"fmt"

	"github.com/bektosh/matrix/services"
)

func main() {
	var (
		operand string
		res services.Matrix
	)
	matrix1 := services.ReadMatrix()
	matrix1.PrintMatrix()
	fmt.Println("What would you like to do next? You can either add(+) or substract(-) matrixes:")
AGAIN:
	fmt.Scan(&operand)

	if operand != "+" && operand != "-" {
		fmt.Println("Sorry, but you entered an invalid operation! Try again:")
		goto AGAIN
	}
	matrix2 := services.ReadMatrix()
	matrix2.PrintMatrix()
	
	if operand == "+" {
		res = matrix1.Add(matrix2)
	}
	fmt.Println("Here is your result:")
	res.PrintMatrix()
}
