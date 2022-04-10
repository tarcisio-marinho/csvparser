package models

import "fmt"

// UNIQUE FIELD ->
type UniqueFields struct {
	Fields map[string]map[string]bool
}

func (fields UniqueFields) AlreadyInserted(value string, fieldName string) bool {

	for uniqFieldName, fieldDict := range fields.Fields {
		if fieldName == uniqFieldName {

			value, found := fieldDict[value]

			if found {
				fmt.Println(value)
				return true
			}

			return false
		}
	}
	return false
}

func CreateUniqueFields(fields []string) UniqueFields {
	uniqueFields := UniqueFields{Fields: make(map[string]map[string]bool)}
	for _, field := range fields {
		uniqueFields.Fields[field] = make(map[string]bool, 0)
	}

	return uniqueFields
}
