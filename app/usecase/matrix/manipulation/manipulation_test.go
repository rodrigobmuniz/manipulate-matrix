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
func TestStringify1x1(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1"})
	expected := "1\n"

	result, _ := manipulation.Stringify(matrix)

	assert.Equal(expected, result, "receive an matrix 1x1 and return a string that exactly represents the received matrix")
}

func TestStringify2x2(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "10"})
	matrix = append(matrix, []string{"2", "20"})
	expected := "1,10\n2,20\n"

	result, _ := manipulation.Stringify(matrix)

	assert.Equal(expected, result, "receive an matrix 2x2 and return a string that exactly represents the received matrix")
}

func TestStringify3x3(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2", "3"})
	matrix = append(matrix, []string{"4", "5", "6"})
	matrix = append(matrix, []string{"7", "8", "9"})
	expected := "1,2,3\n4,5,6\n7,8,9\n"

	result, _ := manipulation.Stringify(matrix)

	assert.Equal(expected, result, "receive an matrix 3x3 and return a string that exactly represents the received matrix")
}

// Invert test
func TestInvert1x1(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1"})

	var expected [][]string
	expected = append(expected, []string{"1"})
	result, _ := manipulation.Invert(matrix)

	assert.Equal(expected, result, "receive an matrix 3x3 and return a string with the values reverted")
}

func TestInvert2x2(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2"})
	matrix = append(matrix, []string{"3", "4"})

	var expected [][]string
	expected = append(expected, []string{"1", "3"})
	expected = append(expected, []string{"2", "4"})

	result, _ := manipulation.Invert(matrix)

	assert.Equal(expected, result, "receive an matrix 2x2 and return a string with the values reverted")
}

func TestInvert3x3(t *testing.T) {
	assert := assert.New(t)

	var matrix [][]string
	matrix = append(matrix, []string{"1", "2", "3"})
	matrix = append(matrix, []string{"4", "5", "6"})
	matrix = append(matrix, []string{"7", "8", "9"})

	var expected [][]string
	expected = append(expected, []string{"1", "4", "7"})
	expected = append(expected, []string{"2", "5", "8"})
	expected = append(expected, []string{"3", "6", "9"})

	result, _ := manipulation.Invert(matrix)

	assert.Equal(expected, result, "receive an matrix 3x3 and return a string with the values reverted")
}
