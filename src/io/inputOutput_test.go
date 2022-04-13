package io

import (
	"csvparser/src/models"
	"reflect"
	"testing"
)

func TestGenerateOutputFiles(t *testing.T) {
	type args struct {
		originalFile string
		employees    []models.Employee
	}
	tests := []struct {
		name string
		args args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateOutputFiles(tt.args.originalFile, tt.args.employees)
		})
	}
}

func TestGetInputFiles(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetInputFiles()
			if got != tt.want {
				t.Errorf("GetInputFiles() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetInputFiles() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestLoadFieldsFromConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want models.Fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadFieldsFromConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadFieldsFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_createOutputDirectory(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createOutputDirectory(); got != tt.want {
				t.Errorf("createOutputDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateOutputFileNames(t *testing.T) {
	tests := []struct {
		name         string
		originalPath string
		want         string
		want1        string
	}{
		{
			name:         "valid filepath",
			originalPath: "/home/tarcisio/file.csv",
			want:         "file-correct.csv",
			want1:        "file-bad.csv",
		},
		{
			name:         "invalid filepath",
			originalPath: "",
			want:         "",
			want1:        "",
		},
		{
			name:         "correct filepath without extension",
			originalPath: "/home/tarcisio/file",
			want:         "file-correct",
			want1:        "file-bad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := generateOutputFileNames(tt.originalPath)
			if got != tt.want {
				t.Errorf("generateOutputFileNames() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("generateOutputFileNames() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getNewCsvHeaders(t *testing.T) {
	type args struct {
		employee models.Employee
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getNewCsvHeaders(tt.args.employee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getNewCsvHeaders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_storeFile(t *testing.T) {
	type args struct {
		filepath  string
		employees []models.Employee
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			storeFile(tt.args.filepath, tt.args.employees)
		})
	}
}
