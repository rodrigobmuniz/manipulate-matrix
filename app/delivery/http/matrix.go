package matrix

import (
	"encoding/csv"
	"fmt"
	"manipulate-matrix/app/usecase/matrix/manipulation"
	"net/http"
	"strings"
)

func HaveError(err error, res http.ResponseWriter) bool {
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return true
	}
	return false
}

func Echo(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if HaveError(err, res) {
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if HaveError(err, res) {
		return
	}
	matrixAsString := manipulation.Stringify(records)
	fmt.Fprint(res, matrixAsString)
}

func Invert(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if HaveError(err, res) {
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if HaveError(err, res) {
		return
	}
	invertedMatrix := manipulation.Invert(records)
	invertedMatrixAsString := manipulation.Stringify(invertedMatrix)
	fmt.Fprint(res, invertedMatrixAsString)
}

func Flatten(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if HaveError(err, res) {
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if HaveError(err, res) {
		return
	}
	flatedMatrix := manipulation.Flatten(records)
	var flatedMatrixAsString string
	flatedMatrixAsString = fmt.Sprintf("%s%s\n", flatedMatrixAsString, strings.Join(flatedMatrix, ","))
	fmt.Fprint(res, flatedMatrixAsString)
}

func Sum(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if HaveError(err, res) {
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if HaveError(err, res) {
		return
	}
	sumOfMatrixItems, _ := manipulation.Sum(records)
	fmt.Fprint(res, sumOfMatrixItems)
}

func Multiply(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if HaveError(err, res) {
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if HaveError(err, res) {
		return
	}
	multiplicationOfMatrixItems, _ := manipulation.Multiply(records)
	fmt.Fprint(res, multiplicationOfMatrixItems)
}
