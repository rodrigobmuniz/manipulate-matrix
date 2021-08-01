package matrix

import (
	"encoding/csv"
	"fmt"
	"manipulate-matrix/app/usecase/matrix/manipulation"
	"net/http"
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
	response, _ := manipulation.PassAlong(records)
	fmt.Fprint(res, response)
}
