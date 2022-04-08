package utils

import (
	"encoding/json"
	"sort"
	"strings"
)

func Pprint(obj interface{}) string {

	fieldsJson, err := json.MarshalIndent(obj, "", " ")

	if err != nil {
		return ""
	}

	return string(fieldsJson)
}

func Trim(s string) string {
	trimmed := strings.ReplaceAll(
		strings.ReplaceAll(
			strings.TrimSpace(strings.ToLower(s)),
			" ", ""),
		"\uFEFF", "")

	return trimmed
}

func SortMapByKey(unsortedMap map[string]int) []int {

	keys := make([]string, 0, len(unsortedMap))

	for k := range unsortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	idxOrder := make([]int, 0, len(unsortedMap))

	for _, k := range keys {
		idxOrder = append(idxOrder, unsortedMap[k])
	}
	return idxOrder
}
