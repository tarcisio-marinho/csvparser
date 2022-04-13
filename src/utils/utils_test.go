package utils

import (
	"reflect"
	"testing"
)

func TestHasValue(t *testing.T) {
	tests := []struct {
		name  string
		value string
		want  bool
	}{
		{
			name:  "empty string",
			value: "",
			want:  false,
		},
		{
			name:  "space string",
			value: " ",
			want:  false,
		},
		{
			name:  "newline string",
			value: "\n",
			want:  false,
		},
		{
			name:  "tab string",
			value: "\t",
			want:  false,
		},
		{
			name:  "tab string",
			value: "testing",
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasValue(tt.value); got != tt.want {
				t.Errorf("HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSortMap(t *testing.T) {
	tests := []struct {
		name        string
		unsortedMap map[string]int
		reference   []string
		want        []int
	}{
		{
			name: "1 testing order",
			unsortedMap: map[string]int{
				"t": 1, "e": 2, "s": 3,
			},
			reference: []string{"e", "t", "s"},
			want:      []int{2, 1, 3},
		},
		{
			name: "1 testing order",
			unsortedMap: map[string]int{
				"teste1": 1, "teste2": 2, "teste3": 3,
			},
			reference: []string{"teste3", "teste2", "teste1"},
			want:      []int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortMap(tt.unsortedMap, tt.reference); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "testing normal string",
			s:    "teste",
			want: "teste",
		},
		{
			name: "testing space string",
			s:    " ",
			want: "",
		},
		{
			name: "testing new line string",
			s:    "\n",
			want: "",
		},
		{
			name: "testing uppercase string",
			s:    "TESTE",
			want: "teste",
		},
		{
			name: "testing encoded string",
			s:    "\uFEFFTESTE",
			want: "teste",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.s); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
