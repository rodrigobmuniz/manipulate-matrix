package manipulation

import (
	"fmt"
	"strings"
)

func PassAlong(records [][]string) (string, error) {
	//add error return for not [][]string
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response, nil
}
