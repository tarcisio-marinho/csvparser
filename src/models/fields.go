package models

import (
	"encoding/json"
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

	return fields, nil
}
