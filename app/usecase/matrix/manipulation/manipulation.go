package manipulation

import (
	"fmt"
	"strconv"
	"strings"
)

// Interface to public methods of matrix manipulation
type MatrixManipulation func([][]string) string

// Return a slice with the matrix flatted
func convertMatrixToSlice(matrix [][]string) []string {
	var flattedMatrix []string
	for key := range matrix {
		flattedMatrix = append(flattedMatrix, matrix[key]...)
	}
	return flattedMatrix
}

// Returns a string with the matrix values
func Stringify(matrix [][]string) string {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

// Return the matrix where the columns and rows are inverted
func Invert(matrix [][]string) string {
	invertMatrix := make([][]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			invertMatrix[j] = append(invertMatrix[j], matrix[i][j])
		}
	}
	return Stringify(invertMatrix)
}

// Return a string with the matrix flatted
func Flatten(matrix [][]string) string {
	flattedMatrix := convertMatrixToSlice(matrix)
	var flatedMatrixAsString string
	flatedMatrixAsString = fmt.Sprintf("%s%s\n", flatedMatrixAsString, strings.Join(flattedMatrix, ","))
	return flatedMatrixAsString
}

// Return the sum of the integers in the matrix
func Sum(matrix [][]string) string {
	var flattedMatrix []string
	total := 0
	flattedMatrix = convertMatrixToSlice(matrix)
	for _, value := range flattedMatrix {
		valueToInt, _ := strconv.Atoi(value)
		total += valueToInt
	}
	return strconv.Itoa(total)
}

// Return the product of the integers in the matrix
func Multiply(matrix [][]string) string {
	var flattedMatrix []string
	total := 0
	flattedMatrix = convertMatrixToSlice(matrix)
	for key, value := range flattedMatrix {
		valueToInt, _ := strconv.Atoi(value)
		if key == 0 {
			total = valueToInt
		}
		if key != 0 {
			total = total * valueToInt
		}
	}
	return strconv.Itoa(total)
}
