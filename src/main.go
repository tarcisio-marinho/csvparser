package main

import (
	"csvparser/src/models"
	"csvparser/src/parser"
	"csvparser/src/utils"
	"log"
	"os"
)

func main() {

	if len(os.Args) != 2 { // TODO: melhorar log fatal
		log.Fatal(`
	Program with wrong usarge, the correct one is:
		~$ go run main.go [csvfilePath]`)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err) // TODO: melhorar log
	}
	defer f.Close()

	// TODO: iterar sobre todos os arquivos dentro de uma pasta ?

	requiredFields := models.CreateFields() // TODO: load required fields from disk
	employees := parser.Parse(f, requiredFields)

	for _, employee := range employees {
		if employee.IsCorrect() {
			utils.Pprint("CORRECT:", employee.Data)

		} else {
			utils.Pprint("WRONG:", employee.Data)
		}
	}
}
