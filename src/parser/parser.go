package parser

import (
	"csvparser/src/models"
	"csvparser/src/utils"
	"encoding/csv"
	"io"
	"log"
	"strings"
)

//TODO: Create godoc for each method
func Parse(f io.Reader, requiredFields models.RequiredFields) []models.Employee {

	csvReader := csv.NewReader(f)
	firstRow, err := csvReader.Read()

	if err != nil {
		log.Fatal(err) // TODO: melhorar log
	}

	headers := getHeadersIdx(firstRow)
	reqFieldsIdxs := getRequiredFieldsIdx(headers, requiredFields)

	employees := generateOutput(csvReader, reqFieldsIdxs)

	return employees
}

func generateOutput(csvReader *csv.Reader, fields []models.FieldIndex) []models.Employee {
	correctEntries := make([]models.Employee, 0)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err) // TODO: melhorar log
		}
		employee := models.CreateEmployee()

		for _, field := range fields {
			if !field.MultipleCol {
				value := row[field.Index[0]]
				employee.Data[field.FieldName] = value
			} else {
				finalValue := ""
				for _, value := range field.Index {
					partial := row[value]
					if utils.HasValue(partial) {
						finalValue += partial + " "
					} else {
						finalValue = "" // if any field is empty, exit
						break
					}
				}

				employee.Data[field.FieldName] = strings.TrimSpace(finalValue)
			}
		}
		correctEntries = append(correctEntries, employee)
	}

	return correctEntries
}

// TODO: e se os campos estiverem vazios ?
func getRequiredFieldsIdx(headers map[string]int, reqFields models.RequiredFields) []models.FieldIndex {

	// TODO: tratar o caso se estiver tudo vazio, não achar nenhum campo dos que são requeridos
	fieldsIdx := make([]models.FieldIndex, 0, len(reqFields.Fields))

	for reqField, fieldInfo := range reqFields.Fields {

		for _, info := range fieldInfo {

			if !info.MultipleCol {

				columnIndex, found := headers[info.Name[0]]
				if found {
					fieldsIdx = append(fieldsIdx, models.FieldIndex{
						FieldName:   reqField,
						Index:       []int{columnIndex},
						MultipleCol: info.MultipleCol,
					})
				}

			} else {

				foundHeaders := 0
				tempHeadersIndexDict := make(map[string]int)

				for _, singleColumn := range info.Name {
					columnIndex, found := headers[singleColumn]
					if found {
						foundHeaders++
						tempHeadersIndexDict[singleColumn] = columnIndex
					}
				}

				if foundHeaders > 0 {
					// sorting the keys, under the assumption that 'f' comes first than 'l' from 'first' 'last' names
					fieldsIdx = append(fieldsIdx, models.FieldIndex{
						FieldName:   reqField,
						Index:       utils.SortMapByKey(tempHeadersIndexDict),
						MultipleCol: info.MultipleCol,
					})
				}
			}
		}
	}

	return fieldsIdx
}

func getHeadersIdx(row []string) map[string]int {

	headers := make(map[string]int)

	for index, field := range row {
		trimmedField := utils.Trim(field)
		headers[trimmedField] = index
	}
	return headers
}
