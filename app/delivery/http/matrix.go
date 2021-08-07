package matrix

import (
	"encoding/csv"
	"fmt"
	"manipulate-matrix/app/helper"
	"manipulate-matrix/app/usecase/matrix/manipulation"
	"net/http"
)

// Handles requests related to matrix manipulation
func processMatrixRequest(res http.ResponseWriter, req *http.Request, fn manipulation.MatrixManipulation) {
	file, _, err := req.FormFile("file")
	if helper.HaveError(err, res) {
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if helper.HaveError(err, res) {
		return
	}
	if helper.CheckIfMatrixIsEmpty(records, res) {
		return
	}
	if !helper.CheckIfMatrixIsSquare(records, res) {
		return
	}
	if !helper.CheckIfAllValuesAreConvertibleToNumber(records, res) {
		return
	}
	funcResult, err := fn(records)
	if helper.HaveError(err, res) {
		return
	}
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
