package main

import (
	"csvparser/src/models"
	"csvparser/src/parser"
	"csvparser/src/utils"
	"fmt"
	"log"
	"os"
)

func main() {
	requiredFields := models.CreateFields()

	if len(os.Args) != 2 {
		log.Fatal(`
Program with wrong usarge, the correct one is:
	~$ go run main.go [csvfilePath]`)
	}

	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	reqFieldsIdxs := parser.Parse(f, requiredFields)
	fmt.Println(utils.Pprint(reqFieldsIdxs))
}
