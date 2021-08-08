package helperHttp_test

import (
	"errors"
	"io"
	"manipulate-matrix/app/delivery/http/helperHttp"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Canary test
func TestCanarySpec(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, true, "Valide if infra for test work")
}

// HaveError test
func TestHaveErrorReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var err error
	var res http.ResponseWriter

	assert.Equal(false, helperHttp.HaveError(err, res), "should return false if error is null")
}

func TestHaveErrorReturningTrue(t *testing.T) {
	assert := assert.New(t)

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	var err = errors.New("an error has occurred")

	assert.Equal(true, helperHttp.HaveError(err, w), "should return true if error have a error")
}

// CheckIfAllValuesAreConvertibleToNumber test
func TestCheckIfAllValuesAreConvertibleToNumberReturningTrue(t *testing.T) {
	assert := assert.New(t)

	var matrixOfStrings [][]string
	matrixOfStrings = append(matrixOfStrings, []string{"1", "10"})
	matrixOfStrings = append(matrixOfStrings, []string{"2", "20"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := helperHttp.CheckIfAllValuesAreConvertibleToNumber(matrixOfStrings, w)

	assert.Equal(true, result, "should matrix an array only and INT and return true")
}

func TestCheckIfAllValuesAreConvertibleToNumberReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var matrixOfStrings [][]string
	matrixOfStrings = append(matrixOfStrings, []string{"1", "10"})
	matrixOfStrings = append(matrixOfStrings, []string{"test", "20"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := helperHttp.CheckIfAllValuesAreConvertibleToNumber(matrixOfStrings, w)

	assert.Equal(false, result, "should receive an matrix only and INT and return true")
}

// CheckIfAllValuesAreConvertibleToNumber test
func TestCheckIfMatrixIsSquareReturningTrue(t *testing.T) {
	assert := assert.New(t)

	var matrixSquare [][]string
	matrixSquare = append(matrixSquare, []string{"1", "10"})
	matrixSquare = append(matrixSquare, []string{"2", "20"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := helperHttp.CheckIfMatrixIsSquare(matrixSquare, w)

	assert.Equal(true, result, "should receive an matrix and return true if it is square")
}

func TesCheckIfMatrixIsSquareReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var matrixNotSquare [][]string
	matrixNotSquare = append(matrixNotSquare, []string{"1", "10", "3"})
	matrixNotSquare = append(matrixNotSquare, []string{"test", "20"})
	matrixNotSquare = append(matrixNotSquare, []string{"1", "10", "3"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := helperHttp.CheckIfMatrixIsSquare(matrixNotSquare, w)

	assert.Equal(false, result, "should receive an matrix only and INT and return true")
}

// Test CheckIfMatrixIsEmpty
func TestCheckIfMatrixIsEmptyReturningTrue(t *testing.T) {
	assert := assert.New(t)

	var matrixEmpty [][]string

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := helperHttp.CheckIfMatrixIsEmpty(matrixEmpty, w)

	assert.Equal(true, result, "should receive an empty matrix and return true")
}

func TestCheckIfMatrix1x1IsEmptyReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"628318530717958647692528676655900576839433879875021164194988918"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := helperHttp.CheckIfMatrixIsEmpty(matrix, w)

	assert.Equal(false, result, "should receive matrix 1x1 and return false")
}

func TestCheckIfMatrix2x2IsEmptyReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "10"})
	matrix = append(matrix, []string{"20", "628318530717958647692528676655900576839433879875021164194988918"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := helperHttp.CheckIfMatrixIsEmpty(matrix, w)

	assert.Equal(false, result, "should receive matrix 2x2 and return false")
}