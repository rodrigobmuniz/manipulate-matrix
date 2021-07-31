package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/echo", func(res http.ResponseWriter, req *http.Request) {
		fmt.Println(req)
		file, _, err := req.FormFile("file")

		if err != nil {
			res.Write([]byte(fmt.Sprintf("error %s", err.Error())))
			return
		}

		fmt.Println(file)

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
		var response string
		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
		fmt.Fprint(res, response)
	})
	http.ListenAndServe(":8080", nil)
}
