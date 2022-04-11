package utils

import (
	"encoding/json"
	"fmt"
	"strings"
)

func Pprint(message string, obj interface{}) {

	fieldsJson, err := json.MarshalIndent(obj, "", " ")

	if err != nil {
		fmt.Println(message)
	}

	fmt.Println(message, string(fieldsJson))
}

func HasValue(value string) bool {
	return len(strings.TrimSpace(value)) != 0
}

func Trim(s string) string {
	trimmed := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.TrimSpace(strings.ToLower(s)),
			" ", ""),
		"\uFEFF", "")

	return trimmed
}

func SortMap(unsortedMap map[string]int, reference []string) []int {

	var newOrder []int
	for _, value := range reference {
		v, found := unsortedMap[value]
		if found {
			newOrder = append(newOrder, v)
		}
	}

	return newOrder
}
