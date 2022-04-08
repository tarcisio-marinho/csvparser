package parser

import (
	"csvparser/src/models"
	"csvparser/src/utils"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"sort"
	"strings"
)

//TODO: Create godoc for each method
func Parse(f io.Reader, requiredFields map[string][]models.Info) map[string][]int {
	csvReader := csv.NewReader(f)
	firstRow, err := csvReader.Read()

	if err != nil {
		log.Fatal(err)
	}

	headers := make(map[string]int)

	for index, field := range firstRow {
		trimmed := strings.ReplaceAll(
			strings.ReplaceAll(
				strings.TrimSpace(strings.ToLower(field)),
				" ", ""),
			"\uFEFF", "")
		headers[trimmed] = index
	}

	fmt.Println("headers:", utils.Pprint(headers))

	reqFieldsIdxs := make(map[string][]int)
	// for i in [email, name, id, salary]
	for requiredField, fieldsInfo := range requiredFields {

		for _, info := range fieldsInfo {

			if !info.MultipleCol {
				// apenas uma coluna ->
				// checar se o name está nos headers

				columnIndex, found := headers[info.Name[0]]
				if found {
					reqFieldsIdxs[requiredField] = []int{columnIndex}
				} // should break the loop ?

			} else {

				foundHeaders := 0
				newHeadersIndexDict := make(map[string]int)

				for _, singleColumn := range info.Name {
					columnIndex, found := headers[singleColumn]
					if found { // resolver o problema da ordenação, como garantir que o first vai ser primeiro que o last
						foundHeaders++
						newHeadersIndexDict[singleColumn] = columnIndex
					}
				}

				if foundHeaders > 0 {
					// sorting the keys, under the assumption that 'f' comes first than 'l' from 'first' 'last' names
					reqFieldsIdxs[requiredField] = sortMapByKey(newHeadersIndexDict)
				}
			}
		}
	}

	return reqFieldsIdxs
}

func sortMapByKey(unsortedMap map[string]int) []int {

	keys := make([]string, 0, len(unsortedMap))

	for k := range unsortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	idxOrder := make([]int, 0, len(unsortedMap))

	for _, k := range keys {
		idxOrder = append(idxOrder, unsortedMap[k])
	}
	return idxOrder
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
