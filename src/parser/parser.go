package parser

import (
	"csvparser/src/models"
	"csvparser/src/utils"
	"csvparser/src/validations"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"strings"
)

// Parse is the method that returns the employees in the csv.
//
// Parse receives the io.Reader and the models.Fields, which are the required fields and unique fields.
//
// Parse will output a list of employees.
func Parse(f io.Reader, fields models.Fields) ([]models.Employee, error) {

	log.Println("Validating required and unique fields")
	err := validations.ValidateFields(fields)
	if err != nil {
		log.Printf("Error validating fields: %s", err)
		return nil, err
	}

	csvReader := csv.NewReader(f)
	firstRow, err := csvReader.Read()

	if err != nil {
		log.Printf("Csv file wihtout header, found: %s", err)
		return nil, err
	}

	log.Println("Parsing the csv headers")
	headers, err := getCsvHeadersIdx(firstRow)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Getting the required fields indexes")
	reqFieldsIdxs, err := getRequiredFieldsIndex(headers, fields.RequiredFields)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("Getting employees list")
	employees, err := getEmployees(csvReader, reqFieldsIdxs, fields.GetUniqueFieldsDict())

	if err != nil {
		return nil, err
	}

	return employees, nil
}

// getRequiredFieldsIndex gets the index in the csv for each required field
//
// given the headers in the csv. Example: name, salary, email -> {"name":0, "salary":1, "email":2}
//
// and the required fields as map[string][]models.Field
//
// outputs a list of models.FieldIndex and error if required fields not found or are incomplete
//
// E.g. requiredFields: ["name", "email"] - csv: firstname, lastname, id, e-mail
//
// output - [ {"name": {"Index": [0, 1], multipleColumn: true}}, {"email": {"Index": [3], multipleColumn: false}} ]
func getRequiredFieldsIndex(headers map[string]int, requiredFields map[string][]models.Field) ([]models.FieldIndex, error) {

	requiredFieldsIdx := make([]models.FieldIndex, 0, len(requiredFields))

	for reqField, fieldData := range requiredFields {

		for _, data := range fieldData {

			if !data.MultipleCol {

				columnIndex, found := headers[data.GetName()]
				if found {
					requiredFieldsIdx = append(requiredFieldsIdx,
						models.FieldIndex{
							FieldName:   reqField,
							Index:       []int{columnIndex},
							MultipleCol: data.MultipleCol,
						})
					break
				}

			} else {

				foundHeaders := 0
				tempHeadersIndexDict := make(map[string]int)

				for _, columnPart := range data.Name {
					columnIndex, found := headers[columnPart]
					if found {
						foundHeaders++
						tempHeadersIndexDict[columnPart] = columnIndex
					}
				}

				if foundHeaders == len(data.Name) {
					requiredFieldsIdx = append(requiredFieldsIdx, models.FieldIndex{
						FieldName:   reqField,
						Index:       utils.SortMap(tempHeadersIndexDict, data.Name),
						MultipleCol: data.MultipleCol,
					})
				}
			}
		}
	}

	if len(requiredFieldsIdx) != len(requiredFields) {
		return nil, errors.New("CSV headers doesn't match with the required fields. Some required fields where not found")
	}

	return requiredFieldsIdx, nil
}

// getEmployees process the csv file content getting the employees required fields
//
// and checks duplicated fields by the configured models.UniqueFields
//
// outputs a list of models.Employee
func getEmployees(csvReader *csv.Reader, fields []models.FieldIndex, uniqueFields models.UniqueFields) ([]models.Employee, error) {
	employees := make([]models.Employee, 0)

	// TODO: testar cenário que a row não tem um dos campos, campos a menos, campos a mais

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("Error reading the csv: %s", err)
			return nil, err
		}

		employee := models.CreateEmployee()

		for _, field := range fields {

			if !field.MultipleCol {
				value := row[field.Index[0]]

				insertEmpAndCheckUniquiness(&employee, uniqueFields, value, field.FieldName)

			} else {
				finalValue := ""
				for _, value := range field.Index {
					partial := row[value]

					if utils.HasValue(partial) {
						finalValue += partial + " "
					}
				}

				insertEmpAndCheckUniquiness(&employee, uniqueFields, finalValue, field.FieldName)
			}
		}
		employees = append(employees, employee)
	}

	return employees, nil
}

// insertEmpAndCheckUniquiness updates the employee data, marks it as correct or incorrect
// regarding if its unique and if the data is valid
func insertEmpAndCheckUniquiness(employee *models.Employee, uniqueFields models.UniqueFields, value string, fieldName string) {
	employee.Data[fieldName] = strings.TrimSpace(value)

	if utils.HasValue(value) {
		employee.SetCorrect()
	} else {
		employee.SetIncorrect(fmt.Sprintf("empty value for field: %s", fieldName))
	}

	if uniqueFields.AlreadyInserted(value, fieldName) {
		employee.SetIncorrect(fmt.Sprintf("%s: %s - duplicated field", fieldName, value))

	} else {
		uniqueFields.InsertField(value, fieldName)
	}
}

// getCsvHeadersIdx returns the field and index in the csv
//
// E.g. csv: name, salary, email - output: {"name":0, "salary":1, "email":2}
func getCsvHeadersIdx(row []string) (map[string]int, error) {

	headers := make(map[string]int)

	for index, field := range row {
		trimmedField := utils.Trim(field)
		if utils.HasValue(trimmedField) {
			headers[trimmedField] = index
		}
	}

	if len(headers) == 0 {
		return headers, errors.New("empty headers in the csv")
	}

	return headers, nil
}
