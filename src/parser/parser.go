package parser

import (
	"csvparser/src/models"
	"csvparser/src/utils"
	"encoding/csv"
	"fmt"
	"io"
	"log"
)

//TODO: Create godoc for each method
func Parse(f io.Reader, requiredFields map[string][]models.Info) map[string][]int {

	csvReader := csv.NewReader(f)
	firstRow, err := csvReader.Read()

	if err != nil {
		log.Fatal(err)
	}

	headers := getHeadersIdx(firstRow)
	reqFieldsIdxs := getRequiredFieldsIdx(headers, requiredFields)

	fmt.Println("headers:", utils.Pprint(headers)) // TODO: remove this print

	return reqFieldsIdxs
}

func getRequiredFieldsIdx(headers map[string]int, requiredFields map[string][]models.Info) map[string][]int {

	reqFieldsIdxs := make(map[string][]int)

	for requiredField, fieldsInfo := range requiredFields {

		for _, info := range fieldsInfo {

			if !info.MultipleCol {

				columnIndex, found := headers[info.Name[0]]
				if found {
					reqFieldsIdxs[requiredField] = []int{columnIndex}
				} // should break the loop ?

			} else {

				foundHeaders := 0
				newHeadersIndexDict := make(map[string]int)

				for _, singleColumn := range info.Name {
					columnIndex, found := headers[singleColumn]
					if found {
						foundHeaders++
						newHeadersIndexDict[singleColumn] = columnIndex
					}
				}

				if foundHeaders > 0 {
					// sorting the keys, under the assumption that 'f' comes first than 'l' from 'first' 'last' names
					reqFieldsIdxs[requiredField] = utils.SortMapByKey(newHeadersIndexDict)
				}
			}
		}
	}

	return reqFieldsIdxs
}

func getHeadersIdx(row []string) map[string]int {

	headers := make(map[string]int)

	for index, field := range row {
		trimmedField := utils.Trim(field)
		headers[trimmedField] = index
	}
	return headers
}

func parserCsv(f io.Reader) {

	csvReader := csv.NewReader(f)
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// do something with read line
		fmt.Printf("%+v\n", row)
	}
}
