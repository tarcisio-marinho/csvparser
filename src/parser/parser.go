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
func Parse(f io.Reader, fields models.Fields) []models.Employee {

	csvReader := csv.NewReader(f)
	firstRow, err := csvReader.Read()

	if err != nil {
		log.Fatal(err) // TODO: melhorar log
	}

	headers := getCsvHeadersIdx(firstRow)

	reqFieldsIdxs := getRequiredFieldsIndex(headers, fields.RequiredFields)

	employees := getEmployeesData(csvReader, reqFieldsIdxs, fields.GetUniqueFieldsDict())

	return employees
}

//TODO: refatorar todas as variaveis desse codigo
func getRequiredFieldsIndex(headers map[string]int, requiredFields map[string][]models.Field) []models.FieldIndex {

	// TODO: tratar o caso se estiver tudo vazio, não achar nenhum campo dos que são requeridos
	fieldsIdx := make([]models.FieldIndex, 0, len(requiredFields))

	for reqField, fieldInfo := range requiredFields {

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
					// TODO: Retirar essa lógica de sorting daqui, colocar no dicionário,
					// levando em consideração a ordem do dicionario
					// ["first", "name"] -> first name ,,,, ["name", "first"] -> name first
					//
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

func getEmployeesData(csvReader *csv.Reader, fields []models.FieldIndex, uniqueFields models.UniqueFields) []models.Employee {
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

				if utils.HasValue(value) {
					employee.Correct = true
				} else {
					employee.Correct = false
				}

				if uniqueFields.AlreadyInserted(value, field.FieldName) {

				}

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

func getCsvHeadersIdx(row []string) map[string]int {

	headers := make(map[string]int)

	for index, field := range row {
		trimmedField := utils.Trim(field)
		headers[trimmedField] = index
	}
	return headers
}
