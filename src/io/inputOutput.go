package io

import (
	"csvparser/src/models"
	"csvparser/src/utils"
	"encoding/csv"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func generateOutputFileNames(originalPath string) (string, string) {
	correct := "-correct"
	wrong := "-bad"

	if !utils.HasValue(originalPath) {
		return "", ""
	}

	filename := filepath.Base(originalPath)

	ext := filepath.Ext(originalPath)

	if utils.HasValue(ext) {
		return strings.ReplaceAll(filename, ext, "") + correct + ext,
			strings.ReplaceAll(filename, ext, "") + wrong + ext
	}

	return filename + correct, filename + wrong

}

func createOutputDirectory() string {
	defaultDirectory := "OUTPUT/"

	if _, err := os.Stat(defaultDirectory); err != nil {
		_ = os.MkdirAll(defaultDirectory, 0777)
	}

	return defaultDirectory
}

func GenerateOutputFiles(originalFile string, employees []models.Employee) {

	correctFileName, wrongFileName := generateOutputFileNames(originalFile)

	if !utils.HasValue(correctFileName) || !utils.HasValue(wrongFileName) {
		log.Fatal("output files are empty")
	}

	directoryPath := createOutputDirectory()

	correctEmployees := make([]models.Employee, 0)
	wrongEmployees := make([]models.Employee, 0)

	for _, employee := range employees {
		if employee.IsCorrect() {
			log.Printf("Correct employee: %s", employee.Data)
			correctEmployees = append(correctEmployees, employee)
		} else {
			log.Printf("Wrong employee: %s, reason: %s", employee.Data, employee.Correct.Reason)
			wrongEmployees = append(wrongEmployees, employee)
		}
	}

	storeFile(directoryPath+wrongFileName, wrongEmployees)
	storeFile(directoryPath+correctFileName, correctEmployees)
}

func storeFile(filepath string, employees []models.Employee) {

	if len(employees) == 0 {
		return
	}

	if _, err := os.Stat(filepath); !errors.Is(err, os.ErrNotExist) {
		log.Fatal("Error generating output, file:", filepath, " already exists")
	}

	csvFile, err := os.Create(filepath)

	if err != nil {
		log.Fatalf("Failed creating file: %s", err)
	}

	csvWriter := csv.NewWriter(csvFile)

	newCsvHeader := getNewCsvHeaders(employees[0])
	csvWriter.Write(newCsvHeader)

	for _, employee := range employees {
		var row []string

		for _, fieldName := range newCsvHeader {
			data, found := employee.Data[fieldName]
			if found {
				row = append(row, data)
			}
		}
		csvWriter.Write(row)
	}

	log.Printf("Output file: %s created", filepath)
	csvWriter.Flush()
	defer csvFile.Close()
}

func getNewCsvHeaders(employee models.Employee) []string {
	keys := make([]string, 0, len(employee.Data))
	for k := range employee.Data {
		keys = append(keys, k)
	}
	return keys
}

func LoadFieldsFromConfig(path string) models.Fields {
	fileContent, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	exampleFormat := `
{
  "RequiredFields": {
    "name": [
      {
        "Name": [
          "name"
        ],
        "MultipleCol": false
      },
      {
        "Name": [
          "f.name",
          "l.name"
        ],
        "MultipleCol": true
      }
    ],
    "salary": [
      {
        "Name": [
          "rate"
        ],
        "MultipleCol": false
      }
    ]
  },
  "UniqueFields": [
      "name"
    ]
}

In this example, the required fields that will be searched in the csv are: "name" and "salary"
with the possibles headers "name", "f.name" + "l.name" and "rate" respectively
if any unique fields, it must be in the required fields
Unique fields cannot be repeated in the csv, it must be unique
	`
	fields, err := models.CreateFieldsFromConfig(fileContent)

	if err != nil {
		log.Fatal("error in the json config format, the correct format should be:\n",
			exampleFormat, "the error: ", err)
	}

	return fields
}

func GetInputFiles() (string, string) {

	if len(os.Args) != 3 {
		log.Fatal(`
			Program with wrong usage, the correct one is:
				~$ go run src/main.go [csvFilePath] [configFilePath]
			example:
				~$ go run src/main.go samples/roster1.csv config/full_config.json
					`)
	}

	return os.Args[1], os.Args[2]
}
