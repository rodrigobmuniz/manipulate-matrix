package matrix_test

import (
	"errors"
	"io"
	matrix "manipulate-matrix/app/delivery/http"
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

	assert.Equal(false, matrix.HaveError(err, res), "should return false if error is null")
}

func TestHaveErrorReturningTrue(t *testing.T) {
	assert := assert.New(t)

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)
	var err = errors.New("an error has occurred")

	assert.Equal(true, matrix.HaveError(err, w), "should return true if error have a error")
}

// AllValuesAreConvertibleToNumber test
func TestAllValuesAreConvertibleToNumberReturningTrue(t *testing.T) {
	assert := assert.New(t)

	var matrixOfStrings [][]string
	matrixOfStrings = append(matrixOfStrings, []string{"1", "10"})
	matrixOfStrings = append(matrixOfStrings, []string{"2", "20"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := matrix.AllValuesAreConvertibleToNumber(matrixOfStrings, w)

	assert.Equal(true, result, "should matrix an array only and INT and return true")
}

func TestAllValuesAreConvertibleToNumberReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var matrixOfStrings [][]string
	matrixOfStrings = append(matrixOfStrings, []string{"1", "10"})
	matrixOfStrings = append(matrixOfStrings, []string{"test", "20"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := matrix.AllValuesAreConvertibleToNumber(matrixOfStrings, w)

	assert.Equal(false, result, "should receive an matrix only and INT and return true")
}

// AllValuesAreConvertibleToNumber test
func TestMatrixIsSquare(t *testing.T) {
	assert := assert.New(t)

	var matrixSquare [][]string
	matrixSquare = append(matrixSquare, []string{"1", "10"})
	matrixSquare = append(matrixSquare, []string{"2", "20"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := matrix.MatrixIsSquare(matrixSquare, w)

	assert.Equal(true, result, "should receive an matrix and return true if it is square")
}

func TesMatrixIsSquareReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var matrixNotSquare [][]string
	matrixNotSquare = append(matrixNotSquare, []string{"1", "10", "3"})
	matrixNotSquare = append(matrixNotSquare, []string{"test", "20"})
	matrixNotSquare = append(matrixNotSquare, []string{"1", "10", "3"})

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "<html><body>Hello World!</body></html>")
	}
	req := httptest.NewRequest("GET", "http://test.com", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	result := matrix.MatrixIsSquare(matrixNotSquare, w)

	assert.Equal(false, result, "should receive an matrix only and INT and return true")
}
