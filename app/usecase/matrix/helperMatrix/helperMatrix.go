package helperMatrix

import (
	"errors"
	"math/big"
)

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
	flattedMatrix := ConvertMatrixToSlice(matrix)
	if len(flattedMatrix) == 0 {
		return true, errors.New("sent matrix is empty")
	}
	return false, nil
}
