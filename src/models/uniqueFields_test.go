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
	type args struct {
		value     string
		fieldName string
	}
	tests := []struct {
		name   string
		Fields map[string]map[string]bool
		args   args
		want   bool
	}{
		{
			name: "already inserted",
			Fields: map[string]map[string]bool{
				"test": {"t": true},
			},
			args: args{
				value:     "e",
				fieldName: "test",
			},
			want: false,
		},
		{
			name: "it was not inserted",
			Fields: map[string]map[string]bool{
				"test": {"t": true},
			},
			args: args{
				value:     "e",
				fieldName: "teste",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := UniqueFields{
				Fields: tt.Fields,
			}
			if got := fields.AlreadyInserted(tt.args.value, tt.args.fieldName); got != tt.want {
				t.Errorf("AlreadyInserted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueFields_InsertField(t *testing.T) {
	type fields1 struct {
	}
	type args struct {
		value     string
		fieldName string
	}
	tests := []struct {
		name   string
		Fields map[string]map[string]bool
		args   args
	}{
		{
			name: "inserting new field",
			Fields: map[string]map[string]bool{
				"test": {"t": true},
			},
			args: args{
				value:     "e",
				fieldName: "test",
			},
		},
		{
			name: "inserting new field",
			Fields: map[string]map[string]bool{
				"test": {"t": true},
			},
			args: args{
				value:     "e",
				fieldName: "teste",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fields := &UniqueFields{
				Fields: tt.Fields,
			}
			fields.InsertField(tt.args.value, tt.args.fieldName)
		})
	}
}
