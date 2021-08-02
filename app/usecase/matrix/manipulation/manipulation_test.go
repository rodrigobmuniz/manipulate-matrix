package manipulation_test

import (
	"manipulate-matrix/app/usecase/matrix/manipulation"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Canary test
func TestCanarySpec(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(true, true, "Valide if infra for test work")
}

// Stringify test
func TestStringify1x1Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1"})
	expected := "1\n"

	result := manipulation.Stringify(matrix)

	assert.Equal(expected, result, "receive an matrix 1x1 and return a string that exactly represents the received matrix")
}

func TestStringify2x2Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "10"})
	matrix = append(matrix, []string{"2", "20"})
	expected := "1,10\n2,20\n"

	result := manipulation.Stringify(matrix)

	assert.Equal(expected, result, "receive an matrix 2x2 and return a string that exactly represents the received matrix")
}

func TestStringify3x3Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2", "3"})
	matrix = append(matrix, []string{"4", "5", "6"})
	matrix = append(matrix, []string{"7", "8", "9"})
	expected := "1,2,3\n4,5,6\n7,8,9\n"

	result := manipulation.Stringify(matrix)

	assert.Equal(expected, result, "receive an matrix 3x3 and return a string that exactly represents the received matrix")
}

// Invert test
func TestInvert1x1Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1"})

	expected := "1\n"
	result := manipulation.Invert(matrix)

	assert.Equal(expected, result, "receive an matrix 3x3 and return a string with the values reverted")
}

func TestInvert2x2Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2"})
	matrix = append(matrix, []string{"3", "4"})

	expected := "1,3\n2,4\n"

	result := manipulation.Invert(matrix)

	assert.Equal(expected, result, "receive an matrix 2x2 and return a string with the values reverted")
}

func TestInvert3x3Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2", "3"})
	matrix = append(matrix, []string{"4", "5", "6"})
	matrix = append(matrix, []string{"7", "8", "9"})

	expected := "1,4,7\n2,5,8\n3,6,9\n"

	result := manipulation.Invert(matrix)

	assert.Equal(expected, result, "receive an matrix 3x3 and return a string with the values reverted")
}

// ConvertMatrixToSlice test
func TestConvertMatrixToSlice2x2(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2"})
	matrix = append(matrix, []string{"3", "4"})

	expected := []string{"1", "2", "3", "4"}

	result := manipulation.ConvertMatrixToSlice(matrix)

	assert.Equal(expected, result, "receive an matrix 3x3 and return a slices with all the numbers in the same order")
}

func TestConvertMatrixToSlice3x3(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2", "3"})
	matrix = append(matrix, []string{"4", "5", "6"})
	matrix = append(matrix, []string{"7", "8", "9"})

	expected := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

	result := manipulation.ConvertMatrixToSlice(matrix)

	assert.Equal(expected, result, "receive an matrix 3x3 and return a slices with all the numbers in the same order")
}

// Flatten test
func TestFlatten1x1Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1"})

	expected := "1"

	result := manipulation.Flatten(matrix)

	assert.Equal(expected, result, "receive an matrix 1x1 and return a slices with all the numbers in the same order")
}

func TestFlatten4x4Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2", "3", "4"})
	matrix = append(matrix, []string{"5", "6", "7", "8"})
	matrix = append(matrix, []string{"9", "10", "11", "12"})
	matrix = append(matrix, []string{"13", "14", "15", "16"})

	expected := "1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16"

	result := manipulation.Flatten(matrix)

	assert.Equal(expected, result, "receive an matrix 4x4 and return a slices with all the numbers in the same order")
}

// Sum test
func TestSum1x1Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"10"})

	expected := "10"

	result := manipulation.Sum(matrix)

	assert.Equal(expected, result, "should return 10 after summing all matrix values")
}

func TestSum2x2Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"10", "20"})
	matrix = append(matrix, []string{"30", "40"})

	expected := "100"

	result := manipulation.Sum(matrix)

	assert.Equal(expected, result, "should return 100 after summing all matrix values")
}

func TestSum3x3Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2", "3"})
	matrix = append(matrix, []string{"-4", "5", "6"})
	matrix = append(matrix, []string{"7", "8", "9"})

	expected := "37"

	result := manipulation.Sum(matrix)

	assert.Equal(expected, result, "should return 45 after summing all matrix values")
}

// Multiply test
func TestMultiply1x1Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"100"})

	expected := "100"

	result := manipulation.Multiply(matrix)

	assert.Equal(expected, result, "should return 100 after multiply all matrix values")
}

func TestMultiply2x2Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"10", "20"})
	matrix = append(matrix, []string{"30", "40"})

	expected := "240000"

	result := manipulation.Multiply(matrix)

	assert.Equal(expected, result, "should return 240000 after multiply all matrix values")
}

func TestMultiply3x3Matrix(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"-1", "20", "10"})
	matrix = append(matrix, []string{"5", "6", "2"})
	matrix = append(matrix, []string{"1", "3", "2"})

	expected := "-72000"

	result := manipulation.Multiply(matrix)

	assert.Equal(expected, result, "should return 72000 after multiply all matrix values")
}
