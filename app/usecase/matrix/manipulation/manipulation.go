package manipulation

import (
	"fmt"
	"strconv"
	"strings"
)

// Returns a string with the matrix values
func Stringify(matrix [][]string) string {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

// Return the matrix where the columns and rows are inverted
func Invert(matrix [][]string) [][]string {
	invertMatrix := make([][]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			invertMatrix[j] = append(invertMatrix[j], matrix[i][j])
		}
	}
	return invertMatrix
}

// Return a slice with the matrix flatted
func Flatten(matrix [][]string) []string {
	var flattedMatrix []string
	for key := range matrix {
		flattedMatrix = append(flattedMatrix, matrix[key]...)
	}
	return flattedMatrix
}

// Return the sum of the integers in the matrix
func Sum(matrix [][]string) (int, error) {
	var flattedMatrix []string
	total := 0
	flattedMatrix = Flatten(matrix)
	for _, value := range flattedMatrix {
		valueToInt, _ := strconv.Atoi(value)
		total += valueToInt
	}
	return total, nil
}

// Return the product of the integers in the matrix
func Multiply(matrix [][]string) (int, error) {
	var flattedMatrix []string
	total := 0
	flattedMatrix = Flatten(matrix)
	for key, value := range flattedMatrix {
		valueToInt, _ := strconv.Atoi(value)
		if key == 0 {
			total = valueToInt
		}
		if key != 0 {
			total = total * valueToInt
		}
	}
	return total, nil
}
