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
	Content [][]int
	Columns int
	Rows    int
}



func TrimEmpties(s []string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == "" {
			fmt.Println("FOUND:", i)
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

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
	fmt.Println("Enter matrix: ")
	var (
		matrix   Matrix
		elements []string
	)

	for {
		var row []int
		buffer := bufio.NewReader(os.Stdin)
		nums, err := buffer.ReadString('\n')
		nums = nums[:len(nums)-1]
		if err != nil {
			log.Fatalf("ERROR: %v", err)
		} else if nums == "" {
			break
		}

		elements = strings.Split(nums, " ")
		elements = TrimEmpties(elements)
		matrix.Columns = len(elements)

		for _, num := range elements {
			number, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("ERROR: %v", err)
			}
			row = append(row, number)
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
