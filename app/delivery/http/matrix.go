package matrix

import (
	"encoding/csv"
	"fmt"
	"manipulate-matrix/app/usecase/matrix/manipulation"
	"net/http"
	"strings"
)

func Echo(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	matrixAsString := manipulation.Stringify(records)
	fmt.Fprint(res, matrixAsString)
}

func Invert(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	invertedMatrix, _ := manipulation.Invert(records)
	invertedMatrixAsString := manipulation.Stringify(invertedMatrix)
	fmt.Fprint(res, invertedMatrixAsString)
}

func Flatten(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	flatedMatrix, _ := manipulation.Flatten(records)
	var flatedMatrixAsString string
	flatedMatrixAsString = fmt.Sprintf("%s%s\n", flatedMatrixAsString, strings.Join(flatedMatrix, ","))
	fmt.Fprint(res, flatedMatrixAsString)
}

func Sum(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	sumOfMatrixItems, _ := manipulation.Sum(records)
	fmt.Fprint(res, sumOfMatrixItems)
}

func Multiply(res http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("file")
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	multiplicationOfMatrixItems, _ := manipulation.Multiply(records)
	fmt.Fprint(res, multiplicationOfMatrixItems)
}
