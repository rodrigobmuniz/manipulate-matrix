package matrix

import (
	"encoding/csv"
	"fmt"
	"manipulate-matrix/app/usecase/matrix/manipulation"
	"net/http"
)

func HaveError(err error, res http.ResponseWriter) bool {
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return true
	}
	return false
}

func ProcessMatrixRequest(res http.ResponseWriter, req *http.Request, fn manipulation.MatrixManipulation) {
	file, _, err := req.FormFile("file")
	if HaveError(err, res) {
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if HaveError(err, res) {
		return
	}
	funcResult := fn(records)
	fmt.Fprint(res, funcResult)
}

func Echo(res http.ResponseWriter, req *http.Request) {
	ProcessMatrixRequest(res, req, manipulation.Stringify)
}

func Invert(res http.ResponseWriter, req *http.Request) {
	ProcessMatrixRequest(res, req, manipulation.Invert)
}

func Flatten(res http.ResponseWriter, req *http.Request) {
	ProcessMatrixRequest(res, req, manipulation.Flatten)
}

func Sum(res http.ResponseWriter, req *http.Request) {
	ProcessMatrixRequest(res, req, manipulation.Sum)
}

func Multiply(res http.ResponseWriter, req *http.Request) {
	ProcessMatrixRequest(res, req, manipulation.Multiply)
}
