package models

import (
	"reflect"
	"testing"
)

func TestCreateUniqueFields(t *testing.T) {
	type args struct {
		fields []string
	}
	tests := []struct {
		name string
		args args
		want UniqueFields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateUniqueFields(tt.args.fields); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUniqueFields() = %v, want %v", got, tt.want)
			}
		})
	}
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
