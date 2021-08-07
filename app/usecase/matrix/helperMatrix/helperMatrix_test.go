package helperMatrix_test

import (
	"manipulate-matrix/app/usecase/matrix/helperMatrix"
	"testing"

	"github.com/stretchr/testify/assert"
)

// AllValuesAreConvertibleToBigInt test
func TestAllValuesAreConvertibleToBigIntReturningTrue(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20", "10"})
	matrix = append(matrix, []string{"5", "6", "2"})
	matrix = append(matrix, []string{"1", "3", "2"})

	expected := true

	result, err := helperMatrix.AllValuesAreConvertibleToBigInt(matrix)

	hasError := err != nil

	assert.Equal(expected, result, "should take an matrix and return true if all items are convertible to int")
	assert.Equal(false, hasError, "should not have generated an Error")
}

func TestAllValuesAreConvertibleToBigIntReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20", "10"})
	matrix = append(matrix, []string{"5", "test", "2"})
	matrix = append(matrix, []string{"1", "3", "2"})

	expected := false

	result, err := helperMatrix.AllValuesAreConvertibleToBigInt(matrix)

	hasError := err != nil

	assert.Equal(expected, result, "should take an matrix and return false if any item are not convertible to int")
	assert.Equal(true, hasError, "should have generated an Error with a message about the problem")
}

// TestTheCheckIfMatrixIsSquare test
func TestTheCheckIfMatrixIsSquare1x1(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1"})

	result, err := helperMatrix.TheCheckIfMatrixIsSquare(matrix)
	hasError := err != nil

	assert.Equal(true, result, "should return true if matrix is square")
	assert.Equal(false, hasError, "should not have generated an Error")
}

func TestTheCheckIfMatrixIsSquare3x3(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20", "10"})
	matrix = append(matrix, []string{"5", "30", "2"})
	matrix = append(matrix, []string{"1", "3", "2"})

	result, err := helperMatrix.TheCheckIfMatrixIsSquare(matrix)
	hasError := err != nil

	assert.Equal(true, result, "should return true if matrix is square")
	assert.Equal(false, hasError, "should not have generated an Error")
}

func TestTheCheckIfMatrixIsSquare5x5(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20", "10", "2", "4"})
	matrix = append(matrix, []string{"5", "30", "2", "1", "90"})
	matrix = append(matrix, []string{"1", "3", "2", "1", "3"})
	matrix = append(matrix, []string{"5", "30", "2", "1", "90"})
	matrix = append(matrix, []string{"1", "3", "2", "1", "3"})

	result, err := helperMatrix.TheCheckIfMatrixIsSquare(matrix)
	hasError := err != nil

	assert.Equal(true, result, "should return true if matrix is square")
	assert.Equal(false, hasError, "should not have generated an Error")
}

func TestTheCheckIfMatrixIsSquare3x3WithError(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20", "10"})
	matrix = append(matrix, []string{"5", "2"})
	matrix = append(matrix, []string{"1", "3", "2"})

	result, err := helperMatrix.TheCheckIfMatrixIsSquare(matrix)
	hasError := err != nil

	assert.Equal(false, result, "should return true if matrix is square")
	assert.Equal(true, hasError, "should have generated an Error")
}

func TestTheCheckIfMatrixIsSquare1x3WithError(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20", "10"})

	result, err := helperMatrix.TheCheckIfMatrixIsSquare(matrix)
	hasError := err != nil

	assert.Equal(false, result, "should return true if matrix is square")
	assert.Equal(true, hasError, "should have generated an Error")
}

func TestTheCheckIfMatrixIsSquare4x2WithError(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20"})
	matrix = append(matrix, []string{"1", "20"})
	matrix = append(matrix, []string{"3", "3"})
	matrix = append(matrix, []string{"-1", "2"})

	result, err := helperMatrix.TheCheckIfMatrixIsSquare(matrix)
	hasError := err != nil

	assert.Equal(false, result, "should return true if matrix is square")
	assert.Equal(true, hasError, "should have generated an Error")
}

// TestTheMatrixIsEmpty test
func TestTheMatrixIsEmptyReturningTrue(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string

	result, err := helperMatrix.TheMatrixIsEmpty(matrix)
	hasError := err != nil

	assert.Equal(true, result, "should return true if matrix is empty")
	assert.Equal(true, hasError, "should have generated an Error")
}

func TestTheMatrixIsEmptyReturningFalse(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20", "10"})
	matrix = append(matrix, []string{"5", "30", "2"})
	matrix = append(matrix, []string{"1", "3", "2"})

	result, err := helperMatrix.TheMatrixIsEmpty(matrix)
	hasError := err != nil

	assert.Equal(false, result, "should return false if matrix is empty")
	assert.Equal(false, hasError, "should not have generated an Error")
}
