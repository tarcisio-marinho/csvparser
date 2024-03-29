package models

type UniqueFields struct {
	Fields map[string]map[string]bool
}

func (fields *UniqueFields) InsertField(value string, fieldName string) {
	for uniqFieldName, _ := range fields.Fields {
		if fieldName == uniqFieldName {
			fields.Fields[fieldName][value] = true
			break
		}
	}
}

func (fields UniqueFields) AlreadyInserted(value string, fieldName string) bool {

	for uniqFieldName, fieldDict := range fields.Fields {
		if fieldName == uniqFieldName {
			_, found := fieldDict[value]

			return found
		}
	}
	return false
}

func CreateUniqueFields(fields []string) UniqueFields {
	uniqueFields := UniqueFields{Fields: make(map[string]map[string]bool)}
	for _, fieldName := range fields {
		uniqueFields.Fields[fieldName] = make(map[string]bool, 0)
	}

	return uniqueFields
}
