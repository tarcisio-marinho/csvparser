package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Fields struct {
	RequiredFields map[string][]Field `json:"RequiredFields"`
	UniqueFields   []string           `json:"UniqueFields"`
}

func (f Fields) GetUniqueFieldsDict() UniqueFields {
	return CreateUniqueFields(f.UniqueFields)
}

func CreateFieldsFromConfig(data []byte) (Fields, error) {

	fields := Fields{}

	err := json.Unmarshal(data, &fields)

	if err != nil {
		return fields, err
	}

	for _, uniqueFieldName := range fields.UniqueFields {
		_, found := fields.RequiredFields[uniqueFieldName]
		if !found {
			return fields, errors.New(
				fmt.Sprintf("unique field: %s not present in required fields",
					uniqueFieldName,
				))
		}
	}

	return fields, nil
}
