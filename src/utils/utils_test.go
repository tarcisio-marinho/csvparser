package utils

import (
	"reflect"
	"testing"
)

func TestHasValue(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasValue(tt.args.value); got != tt.want {
				t.Errorf("HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPprint(t *testing.T) {
	type args struct {
		message string
		obj     interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Pprint(tt.args.message, tt.args.obj)
		})
	}
}

func TestSortMap(t *testing.T) {
	type args struct {
		unsortedMap map[string]int
		reference   []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SortMap(tt.args.unsortedMap, tt.args.reference); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrim(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Trim(tt.args.s); got != tt.want {
				t.Errorf("Trim() = %v, want %v", got, tt.want)
			}
		})
	}
}
