package validations

import (
	"csvparser/src/models"
	"errors"
	"fmt"
)

func ValidateFields(fields models.Fields) error {
	uniqueFields := fields.UniqueFields
	requiredFields := fields.RequiredFields

	if requiredFields == nil || len(requiredFields) == 0 {
		return errors.New("required fields are empty")
	}

	if uniqueFields == nil || len(uniqueFields) == 0 {
		return errors.New("unique fields are empty")
	}

	for fieldName, fieldsData := range requiredFields {

		if fieldsData == nil || len(fieldsData) == 0 {
			return errors.New(
				fmt.Sprintf("required field: %s not configured properly",
					fieldName,
				))
		}

		for _, field := range fieldsData {
			if !field.IsValid() {
				return errors.New(
					fmt.Sprintf("required field: %s - \"%s\" not configured properly",
						fieldName, field.GetName(),
					))
			}
		}
	}

	for _, uniqueFieldName := range uniqueFields {
		_, found := fields.RequiredFields[uniqueFieldName]
		if !found {
			return errors.New(
				fmt.Sprintf("unique field: %s not present in required fields",
					uniqueFieldName,
				))
		}
	}

	return nil
}
