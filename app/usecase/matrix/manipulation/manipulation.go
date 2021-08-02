package manipulation

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Interface to public methods of matrix manipulation
type MatrixManipulation func([][]string) string

// Return a slice with the matrix flatted
func ConvertMatrixToSlice(matrix [][]string) []string {
	var flattedMatrix []string
	for key := range matrix {
		flattedMatrix = append(flattedMatrix, matrix[key]...)
	}
	return flattedMatrix
}

// Return the matrix as a string in matrix format
func Stringify(matrix [][]string) string {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

// Return the matrix as a string in matrix format where the columns and rows are inverted
func Invert(matrix [][]string) string {
	invertMatrix := make([][]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			invertMatrix[j] = append(invertMatrix[j], matrix[i][j])
		}
	}
	return Stringify(invertMatrix)
}

// Return the matrix as a 1 line string, with values separated by commas
func Flatten(matrix [][]string) string {
	flattedMatrix := ConvertMatrixToSlice(matrix)
	var flatedMatrixAsString string
	flatedMatrixAsString = fmt.Sprintf("%s%s", flatedMatrixAsString, strings.Join(flattedMatrix, ","))
	return flatedMatrixAsString
}

// Return the sum of the integers in the matrix
func Sum(matrix [][]string) string {
	var flattedMatrix []string
	total := 0
	flattedMatrix = ConvertMatrixToSlice(matrix)
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
	flattedMatrix = ConvertMatrixToSlice(matrix)
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

func AllValuesAreConvertibleToInt(matrix [][]string) (bool, error) {
	matrixAsSlice := ConvertMatrixToSlice(matrix)
	for _, item := range matrixAsSlice {
		if _, err := strconv.Atoi(item); err != nil {
			errorResponse := fmt.Sprintf("not all items in the matrix are numbers. Check: %s", err.Error())
			var matrixWithWrongContent = errors.New(errorResponse)
			return false, matrixWithWrongContent
		}
	}
	return true, nil
}

func TheMatrixIsSquare(matrix [][]string) (bool, error) {
	totalCols := len(matrix)
	for _, row := range matrix {
		if len(row) != totalCols {
			return false, errors.New("matrix received is not square")
		}
	}
	return true, nil
}
