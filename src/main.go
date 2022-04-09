package main

import (
	"csvparser/src/models"
	"csvparser/src/parser"
	"csvparser/src/utils"
	"encoding/json"
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

	//return "/home/tarcisio/Documents/rain/csvparser/samples/roster1.csv", "config/config.json"
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
	reqFields := loadConfig(configFilePath)

	employees := parser.Parse(f, reqFields)

	for _, employee := range employees {
		if employee.IsCorrect() {
			utils.Pprint("CORRECT:", employee.Data)

		} else {
			utils.Pprint("WRONG:", employee.Data)
		}
	}
}

// TODO: validar arquivo de input -> validar se os arrays estão preenchidos
func loadConfig(path string) models.RequiredFields {
	file, err := ioutil.ReadFile(path)

	if err != nil {
		log.Fatal("error loading config file:", err)
	}

	requiredFields := models.RequiredFields{}

	err = json.Unmarshal([]byte(file), &requiredFields)
	exampleFormat := `
{
  "Fields": {
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
  }
}

In this example, the fields that will be searched are: "name" and "salary"
with the possibles headers "name", "f.name" + "l.name" and "rate" respectively
	`

	if err != nil {
		log.Fatal("error in the json config format, the correct format should be:\n",
			exampleFormat, "the error: ", err)
	}

	return requiredFields

}
