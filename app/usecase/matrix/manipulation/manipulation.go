package manipulation

import (
	"fmt"
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
