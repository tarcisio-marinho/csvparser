package main

import (
	"csvparser/src/models"
	"csvparser/src/parser"
	"csvparser/src/utils"
	"io/ioutil"
	"log"
	"os"
)

func getInputFiles() (string, string) {

	if len(os.Args) != 3 { // TODO: melhorar log fatal
		log.Fatal(`
			Program with wrong usarge, the correct one is:
				~$ go run src/main.go [csvFilePath] [configFilePath]
			example:
				~$ go run src/main.go samples/roster1.csv config/required_fields.json
			`)
	}
	return os.Args[1], os.Args[2]

	/*
		return "/home/tarcisio/Documents/rain/csvparser/samples/roster1.csv",
			"/home/tarcisio/Documents/rain/csvparser/config/full_config.json"
	*/
}

func main() {

	csvFilePath, configFilePath := getInputFiles()

	f, err := os.Open(csvFilePath)

	if err != nil {
		log.Fatal(err) // TODO: melhorar log
	}
	defer f.Close()

	// TODO: iterar sobre todos os arquivos dentro de uma pasta ?

	//requiredFields := models.CreateFields() // TODO: load required fields from disk
	fields := loadFieldsFromConfig(configFilePath)

	employees := parser.Parse(f, fields)

	for _, employee := range employees {
		utils.Pprint("\n\n", employee)
		if employee.IsCorrect() {

			// TODO: implement logic to put on valid output
		} else {
			// TODO: implement logic to put on invalid output
		}
	}
}

// TODO: validar arquivo de input -> validar se os arrays estão preenchidos
func loadFieldsFromConfig(path string) models.Fields {
	fileContent, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("error loading config file:", err)
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
