package services

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Matrix struct {
	Content [][]float32
	Columns int
	Rows    int
}

// func IsNumeric(s string) bool {
// 	_, err := strconv.ParseFloat(s, 32)
//     return err == nil
// }

// IsValid is a method that validates entered matrix. It returns false when entered matrix has wrong structure, otherwise true
func (m *Matrix) IsValid() bool {
	length := len(m.Content[0])

	for i := 1; i < len(m.Content); i++ {
		if len(m.Content[i]) != length {
			return false
		}
	}
	return true
}

// ReadMatrix is a function that reads(scans) a matrix from input.
func ReadMatrix() Matrix {
AGAIN:
	fmt.Println("Enter matrix: ")
	var (
		matrix   Matrix
		elements []string
	)

	for {
		var row []float32
		buffer := bufio.NewReader(os.Stdin)
		nums, err := buffer.ReadString('\n')
		nums = nums[:len(nums)-1]
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		} else if nums == "" {
			break
		}

		elements = strings.Fields(nums)
		matrix.Columns = len(elements)

		for _, num := range elements {
			number, err := strconv.ParseFloat(num, 32)
			if err != nil {
				fmt.Println("Sorry, but seems like your input contains non-numeric values\nPlease check and enter one more time!")
				goto AGAIN
			}
			row = append(row, float32(number))
		}
		matrix.Content = append(matrix.Content, row)
		matrix.Rows++
	}

	if !matrix.IsValid() {
		fmt.Println("The matrix you entered is not valid! Number of columns is not consistent!\nTry again:")
		matrix = ReadMatrix()
	}

	return matrix
}

// PrintMatrix is a method for printing out a matrix in a more accurate way.
// In the form of a mathematic matrixes
func (m *Matrix) PrintMatrix() {
	for _, row := range m.Content {
		fmt.Println(row)
	}
}
