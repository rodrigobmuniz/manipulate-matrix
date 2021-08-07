package manipulation

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

// Interface to public methods of matrix manipulation
type MatrixManipulation func([][]string) (string, error)

// Return the matrix as a string in matrix format
func Stringify(matrix [][]string) (string, error) {
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response, nil
}

// Return the matrix as a string in matrix format where the columns and rows are inverted
func Invert(matrix [][]string) (string, error) {
	invertMatrix := make([][]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			invertMatrix[j] = append(invertMatrix[j], matrix[i][j])
		}
	}
	return Stringify(invertMatrix)
}

// Return the matrix as a 1 line string, with values separated by commas
func Flatten(matrix [][]string) (string, error) {
	flattedMatrix := ConvertMatrixToSlice(matrix)
	var flatedMatrixAsString string
	flatedMatrixAsString = fmt.Sprintf("%s%s", flatedMatrixAsString, strings.Join(flattedMatrix, ","))
	return flatedMatrixAsString, nil
}

// Return the sum of the integers in the matrix
func Sum(matrix [][]string) (string, error) {
	var flattedMatrix []string
	if len(flattedMatrix) > 0 {
		return "", errors.New("empty matrix cannot be summed")
	}
	total := new(big.Int)
	var convertable bool
	total, convertable = total.SetString("0", 10)
	if !convertable {
		return "", errors.New("error on function Sum converting value to bitInt")
	}
	flattedMatrix = ConvertMatrixToSlice(matrix)
	for _, value := range flattedMatrix {
		bitIntValue := new(big.Int)
		valueToInt, _ := bitIntValue.SetString(value, 10)
		total.Add(total, valueToInt)
	}

	return total.String(), nil
}

// Return the product of the integers in the matrix
func Multiply(matrix [][]string) (string, error) {
	var flattedMatrix []string
	if len(flattedMatrix) > 0 {
		return "", errors.New("empty matrix cannot be multiplied")
	}
	total := new(big.Int)
	total, _ = total.SetString("0", 10)
	flattedMatrix = ConvertMatrixToSlice(matrix)
	for key, value := range flattedMatrix {
		bitIntValue := new(big.Int)
		valueToInt, _ := bitIntValue.SetString(value, 10)
		if key == 0 {
			total = valueToInt
		}
		if key != 0 {
			total.Mul(total, valueToInt)
		}
	}
	return total.String(), nil
}

// Return a slice with the matrix flatted
func ConvertMatrixToSlice(matrix [][]string) []string {
	var flattedMatrix []string
	for key := range matrix {
		flattedMatrix = append(flattedMatrix, matrix[key]...)
	}
	return flattedMatrix
}

// Return true if all items are convertible to bitInt or return false, with error, if is not
func AllValuesAreConvertibleToBigInt(matrix [][]string) (bool, error) {
	matrixAsSlice := ConvertMatrixToSlice(matrix)
	for _, item := range matrixAsSlice {
		bitIntValue := new(big.Int)
		if _, converted := bitIntValue.SetString(item, 10); !converted {
			var matrixWithWrongContent = errors.New("not all items in the matrix are numbers")
			return false, matrixWithWrongContent
		}
	}
	return true, nil
}

// Return true if the matrix is square or return false, with error, if is not
func TheCheckIfMatrixIsSquare(matrix [][]string) (bool, error) {
	totalCols := len(matrix)
	for _, row := range matrix {
		if len(row) != totalCols {
			return false, errors.New("matrix received is not square")
		}
	}
	return true, nil
}

// Return true, with error, if matrix is empty or return false if is not empty
func TheMatrixIsEmpty(matrix [][]string) (bool, error) {
	flattedMatrix, _ := Flatten(matrix)
	if flattedMatrix == "" {
		return true, errors.New("sent matrix is empty")
	}
	return false, nil
}
