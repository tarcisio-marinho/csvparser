package models

import (
	"reflect"
	"testing"
)

func TestCreateUniqueFields(t *testing.T) {
	tests := []struct {
		name   string
		fields []string
		want   UniqueFields
	}{
		{
			name:   "two unique fields",
			fields: []string{"field1", "field2"},
			want:   UniqueFields{Fields: uniqueFieldsMock([]string{"field1", "field2"})},
		},
		{
			name:   "empty unique fields",
			fields: nil,
			want:   UniqueFields{Fields: uniqueFieldsMock(nil)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateUniqueFields(tt.fields)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUniqueFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
func uniqueFieldsMock(fields []string) map[string]map[string]bool {
	uniqueFields := make(map[string]map[string]bool)
	for _, field := range fields {
		uniqueFields[field] = make(map[string]bool, 0)
	}
	return uniqueFields
}

func TestUniqueFields_AlreadyInserted(t *testing.T) {
	type fields1 struct {
		Fields map[string]map[string]bool
	}
	type args struct {
		value     string
		fieldName string
	}
	tests := []struct {
		name   string
		fields fields1
		args   args
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := UniqueFields{
				Fields: tt.fields.Fields,
			}
			if got := fields.AlreadyInserted(tt.args.value, tt.args.fieldName); got != tt.want {
				t.Errorf("AlreadyInserted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueFields_InsertField(t *testing.T) {
	type fields1 struct {
		Fields map[string]map[string]bool
	}
	type args struct {
		value     string
		fieldName string
	}
	tests := []struct {
		name   string
		fields fields1
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := &UniqueFields{
				Fields: tt.fields.Fields,
			}
			fields.InsertField(tt.args.value, tt.args.fieldName)
		})
	}
}
