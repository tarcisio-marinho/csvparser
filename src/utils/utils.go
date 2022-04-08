package utils

import "encoding/json"

func Pprint(obj interface{}) string {

	fieldsJson, err := json.MarshalIndent(obj, "", " ")

	if err != nil {
		return ""
	}

	return string(fieldsJson)
}
