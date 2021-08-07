package helper

import (
	"fmt"
	"manipulate-matrix/app/usecase/matrix/manipulation"
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
func AllValuesAreConvertibleToNumber(matrix [][]string, res http.ResponseWriter) bool {
	isConvertible, err := manipulation.AllValuesAreConvertibleToBigInt(matrix)
	if !isConvertible {
		HaveError(err, res)
		return false
	}
	return true
}

// Takes a string matrix and checks if the matrix is square
func CheckIfMatrixIsSquare(matrix [][]string, res http.ResponseWriter) bool {
	isSquare, err := manipulation.TheCheckIfMatrixIsSquare(matrix)
	if !isSquare {
		HaveError(err, res)
		return false
	}
	return true
}

func CheckIfMatrixIsEmpty(matrix [][]string, res http.ResponseWriter) bool {
	isEmpty, err := manipulation.TheMatrixIsEmpty(matrix)
	if isEmpty {
		HaveError(err, res)
		return true
	}
	return false
}
