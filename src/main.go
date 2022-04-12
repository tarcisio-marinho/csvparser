package main

import (
	"csvparser/src/io"
	"csvparser/src/parser"
	"log"
	"os"
)

func main() {

	csvFilePath, configFilePath := io.GetInputFiles()

	f, err := os.Open(csvFilePath)

	if err != nil {
		log.Fatalf("Error opening the file: %s", err)
	}
	defer f.Close()

	fields := io.LoadFieldsFromConfig(configFilePath)

	employees, err := parser.Parse(f, fields)

	if err != nil {
		log.Printf("Parser error: %s", err)
		return
	}

	io.GenerateOutputFiles(csvFilePath, employees)

	log.Println("DONE")
}
