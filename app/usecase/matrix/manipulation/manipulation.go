package manipulation

import (
	"fmt"
	"strconv"
	"strings"
)

func Stringify(matrix [][]string) (string, error) {
	//add error return for not [][]string
	var response string
	for _, row := range matrix {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response, nil
}

func Invert(matrix [][]string) ([][]string, error) {
	invertMatrix := make([][]string, len(matrix))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			invertMatrix[j] = append(invertMatrix[j], matrix[i][j])
		}
	}
	return invertMatrix, nil
}

func Flatten(matrix [][]string) ([]string, error) {
	var flattedMatrix []string
	for key, _ := range matrix {
		flattedMatrix = append(flattedMatrix, matrix[key]...)
	}
	return flattedMatrix, nil
}

func Sum(matrix [][]string) (int, error) {
	var flattedMatrix []string
	total := 0
	flattedMatrix, _ = Flatten(matrix)
	for _, value := range flattedMatrix {
		valueToInt, _ := strconv.Atoi(value)
		total += valueToInt
	}
	return total, nil
}

func Multiply(matrix [][]string) (int, error) {
	var flattedMatrix []string
	total := 0
	flattedMatrix, _ = Flatten(matrix)
	for key, value := range flattedMatrix {
		valueToInt, _ := strconv.Atoi(value)
		if key == 0 {
			total = valueToInt
		}
		if key != 0 {
			total = total * valueToInt
		}
	}
	return total, nil
}
