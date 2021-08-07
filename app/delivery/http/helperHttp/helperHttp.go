package helperHttp

import (
	"fmt"
	"manipulate-matrix/app/usecase/matrix/helperMatrix"
	"net/http"
)

// If you have an error, return the error to the endpoint
func HaveError(err error, res http.ResponseWriter) bool {
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return true
	}
	return false
}

// Takes a string matrix and checks if all values are convertible to number
func CheckIfAllValuesAreConvertibleToNumber(matrix [][]string, res http.ResponseWriter) bool {
	isConvertible, err := helperMatrix.AllValuesAreConvertibleToBigInt(matrix)
	if !isConvertible {
		HaveError(err, res)
		return false
	}
	return true
}

// Takes a string matrix and checks if the matrix is square
func CheckIfMatrixIsSquare(matrix [][]string, res http.ResponseWriter) bool {
	isSquare, err := helperMatrix.TheCheckIfMatrixIsSquare(matrix)
	if !isSquare {
		HaveError(err, res)
		return false
	}
	return true
}

// Takes a string matrix and checks if the matrix is empty
func CheckIfMatrixIsEmpty(matrix [][]string, res http.ResponseWriter) bool {
	isEmpty, err := helperMatrix.TheMatrixIsEmpty(matrix)
	if isEmpty {
		HaveError(err, res)
		return true
	}
	return false
}
