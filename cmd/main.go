package main

import (
	"fmt"

	"github.com/bektosh/matrix/services"
)

func main() {
	var (
		operand string
		res services.Matrix
		err error
	)
	matrix1 := services.ReadMatrix()
	matrix1.PrintMatrix()
	fmt.Println("What would you like to do next? You can either add(+), substract(-) or multiply(*) matrixes:")
AGAIN:
	fmt.Scan(&operand)

	if operand != "+" && operand != "-" && operand != "*" {
		fmt.Println("Sorry, but you entered an invalid operation! Try again:")
		goto AGAIN
	}
	matrix2 := services.ReadMatrix()
	matrix2.PrintMatrix()
	
	if operand == "+" {
		res, err = matrix1.Add(matrix2)
		if err != nil {
			fmt.Println("ERROR! Entered matrix is not compatible for addition with the first matrix")
		}
	} else if operand == "-" {
		res, err = matrix1.Substract(matrix2)
		if err != nil {
			fmt.Println("ERROR! Entered matrix is not compatible for substraction with the first matrix")
		}
	} else {
		res, err = matrix1.Multiply(matrix2)
		if err != nil {
			fmt.Println("ERROR! Entered matrix is not compatible for multiplying with the first matrix")
		}
	}
	fmt.Println("Here is your result:")
	res.PrintMatrix()
}
