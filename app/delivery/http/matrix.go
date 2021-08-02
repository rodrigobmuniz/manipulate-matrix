package matrix

import (
	"encoding/csv"
	"fmt"
	"manipulate-matrix/app/usecase/matrix/manipulation"
	"net/http"
)

// If you have an error, return the error to the endpoint
func HaveError(err error, res http.ResponseWriter) bool {
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return true
	}
	return false
}

// Takes a string matrix and checks if all values are convertible to number
func AllValuesAreConvertibleToNumber(matrix [][]string, res http.ResponseWriter) bool {
	isConvertible, err := manipulation.AllValuesAreConvertibleToInt(matrix)
	if !isConvertible {
		HaveError(err, res)
		return false
	}
	return true
}

func MatrixIsSquare(matrix [][]string, res http.ResponseWriter) bool {
	isSquare, err := manipulation.TheMatrixIsSquare(matrix)
	if !isSquare {
		HaveError(err, res)
		return false
	}
	return true
}

// Handles requests related to matrix manipulation
func processMatrixRequest(res http.ResponseWriter, req *http.Request, fn manipulation.MatrixManipulation) {
	file, _, err := req.FormFile("file")
	if HaveError(err, res) {
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if HaveError(err, res) {
		return
	}
	if !MatrixIsSquare(records, res) {
		return
	}
	if !AllValuesAreConvertibleToNumber(records, res) {
		return
	}
	funcResult := fn(records)
	fmt.Fprint(res, funcResult)
}

// Handles requests to the endpoint /echo, returning the matrix as a string in matrix format
func Echo(res http.ResponseWriter, req *http.Request) {
	processMatrixRequest(res, req, manipulation.Stringify)
}

// Handles requests to the endpoint /invert, returning the matrix as a string in matrix format where the columns and rows are inverted
func Invert(res http.ResponseWriter, req *http.Request) {
	processMatrixRequest(res, req, manipulation.Invert)
}

// Handles requests to the endpoint /flatten, returning the matrix as a 1 line string, with values separated by commas
func Flatten(res http.ResponseWriter, req *http.Request) {
	processMatrixRequest(res, req, manipulation.Flatten)
}

// Handles requests to the endpoint /sum, returning the sum of the integers in the matrix
func Sum(res http.ResponseWriter, req *http.Request) {
	processMatrixRequest(res, req, manipulation.Sum)
}

// Handles requests to the endpoint /multiply, returning the product of the integers in the matrix
func Multiply(res http.ResponseWriter, req *http.Request) {
	processMatrixRequest(res, req, manipulation.Multiply)
}
