package parser

import (
	"bytes"
	"csvparser/src/models"
	"encoding/csv"
	"io"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		f      io.Reader
		fields models.Fields
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Employee
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.f, tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getCsvHeadersIdx(t *testing.T) {
	tests := []struct {
		name    string
		row     []string
		want    map[string]int
		wantErr bool
	}{
		{
			name: "valid csv with 3 fields",
			row:  []string{"name", "salary", "email"},
			want: map[string]int{
				"name":   0,
				"salary": 1,
				"email":  2,
			},
			wantErr: false,
		},
		{
			name: "empty field in csv with 3 fields",
			row:  []string{"", "salary", "email"},
			want: map[string]int{
				"salary": 1,
				"email":  2,
			},
			wantErr: false,
		},
		{
			name:    "all empty fields",
			row:     []string{"", "", ""},
			want:    map[string]int{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getCsvHeadersIdx(tt.row)
			if (err != nil) != tt.wantErr {
				t.Errorf("getCsvHeadersIdx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCsvHeadersIdx() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getEmployees(t *testing.T) {
	type args struct {
		csvReader    *csv.Reader
		fields       []models.FieldIndex
		uniqueFields models.UniqueFields
	}
	tests := []struct {
		name    string
		args    args
		want    []models.Employee
		wantErr bool
	}{
		{
			name: "2 valid employees",
			args: args{
				csvReader: csv.NewReader(
					bytes.NewBuffer([]byte(`
John Doe,doe@test.com,$10.00,1
Mary Jane,Mary@tes.com,$15,2`)),
				),
				fields: []models.FieldIndex{
					{
						FieldName:   "email",
						Index:       []int{1},
						MultipleCol: false,
					},
					{
						FieldName:   "id",
						Index:       []int{3},
						MultipleCol: false,
					},
				},
				uniqueFields: models.UniqueFields{},
			},
			want: []models.Employee{
				{
					Data: map[string]string{
						"email": "doe@test.com",
						"id":    "1",
					},
					Correct: models.CorrectData{
						IsCorrect: true,
						Reason:    "",
					},
				},
				{
					Data: map[string]string{
						"email": "Mary@tes.com",
						"id":    "2",
					},
					Correct: models.CorrectData{
						IsCorrect: true,
						Reason:    "",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "1 valid employee 1 bad employee",
			args: args{
				csvReader: csv.NewReader(
					bytes.NewBuffer([]byte(`
John Doe,doe@test.com,$10.00,1
Mary Jane,,$15,2`)),
				),
				fields: []models.FieldIndex{
					{
						FieldName:   "email",
						Index:       []int{1},
						MultipleCol: false,
					},
					{
						FieldName:   "id",
						Index:       []int{3},
						MultipleCol: false,
					},
				},
				uniqueFields: models.UniqueFields{},
			},
			want: []models.Employee{
				{
					Data: map[string]string{
						"email": "doe@test.com",
						"id":    "1",
					},
					Correct: models.CorrectData{
						IsCorrect: true,
						Reason:    "",
					},
				},
				{
					Data: map[string]string{
						"email": "",
						"id":    "2",
					},
					Correct: models.CorrectData{
						IsCorrect: false,
						Reason:    "empty value for field: email",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "duplicated email field",
			args: args{
				csvReader: csv.NewReader(
					bytes.NewBuffer([]byte(`
John Doe,doe@test.com,$10.00,1
Mary Jane,doe@test.com,$15,2`)),
				),
				fields: []models.FieldIndex{
					{
						FieldName:   "email",
						Index:       []int{1},
						MultipleCol: false,
					},
					{
						FieldName:   "id",
						Index:       []int{3},
						MultipleCol: false,
					},
				},
				uniqueFields: models.UniqueFields{
					Fields: map[string]map[string]bool{
						"email": {"doe@test.com": true},
					},
				},
			},
			want: []models.Employee{
				{
					Data: map[string]string{
						"email": "doe@test.com",
						"id":    "1",
					},
					Correct: models.CorrectData{
						IsCorrect: false,
						Reason:    "email: doe@test.com - duplicated field",
					},
				},
				{
					Data: map[string]string{
						"email": "doe@test.com",
						"id":    "2",
					},
					Correct: models.CorrectData{
						IsCorrect: false,
						Reason:    "email: doe@test.com - duplicated field",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "all empty fields",
			args: args{
				csvReader: csv.NewReader(
					bytes.NewBuffer([]byte(`,,`)),
				),
				fields: []models.FieldIndex{
					{
						FieldName:   "email",
						Index:       []int{1},
						MultipleCol: false,
					},
					{
						FieldName:   "id",
						Index:       []int{2},
						MultipleCol: false,
					},
				},
				uniqueFields: models.UniqueFields{
					Fields: map[string]map[string]bool{
						"email": {"doe@test.com": true},
					},
				},
			},
			want: []models.Employee{
				{
					Data: map[string]string{
						"email": "",
						"id":    "",
					},
					Correct: models.CorrectData{
						IsCorrect: false,
						Reason:    "empty value for field: id",
					},
				},
			},
			wantErr: false,
		},
		{
			name: "valid multi column",
			args: args{
				csvReader: csv.NewReader(
					bytes.NewBuffer([]byte(`
John,Doe,doe@test.com,$10.00,1`)),
				),
				fields: []models.FieldIndex{
					{
						FieldName:   "name",
						Index:       []int{0, 1},
						MultipleCol: true,
					},
					{
						FieldName:   "email",
						Index:       []int{2},
						MultipleCol: false,
					},
					{
						FieldName:   "id",
						Index:       []int{4},
						MultipleCol: false,
					},
				},
				uniqueFields: models.UniqueFields{},
			},
			want: []models.Employee{
				{
					Data: map[string]string{
						"name":  "John Doe",
						"email": "doe@test.com",
						"id":    "1",
					},
					Correct: models.CorrectData{
						IsCorrect: true,
						Reason:    "",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getEmployees(tt.args.csvReader, tt.args.fields, tt.args.uniqueFields)
			if (err != nil) != tt.wantErr {
				t.Errorf("getEmployees() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getEmployees() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getRequiredFieldsIndex(t *testing.T) {
	tests := []struct {
		name           string
		headers        map[string]int
		requiredFields map[string][]models.Field
		want           []models.FieldIndex
		wantErr        bool
	}{
		{
			name:    "valid headers and requried fields",
			headers: map[string]int{"name": 0, "email": 1, "id": 2},
			requiredFields: map[string][]models.Field{
				"email": {
					{
						Name:        []string{"email"},
						MultipleCol: false,
					},
				},

				"id": {
					{
						Name:        []string{"id"},
						MultipleCol: false,
					},
				},
			},

			want: []models.FieldIndex{
				{
					FieldName:   "email",
					Index:       []int{1},
					MultipleCol: false,
				},
				{
					FieldName:   "id",
					Index:       []int{2},
					MultipleCol: false,
				},
			},
			wantErr: false,
		},
		{
			name:    "valid headers and requried fields multi column",
			headers: map[string]int{"fname": 0, "lname": 1, "id": 2},
			requiredFields: map[string][]models.Field{
				"name": {
					{
						Name:        []string{"fname", "lname"},
						MultipleCol: true,
					},
				},

				"id": {
					{
						Name:        []string{"id"},
						MultipleCol: false,
					},
				},
			},

			want: []models.FieldIndex{
				{
					FieldName:   "name",
					Index:       []int{0, 1},
					MultipleCol: true,
				},
				{
					FieldName:   "id",
					Index:       []int{2},
					MultipleCol: false,
				},
			},
			wantErr: false,
		},
		{
			name:    "valid headers and requried fields",
			headers: map[string]int{"asd": 0, "fds": 1, "adsf": 2},
			requiredFields: map[string][]models.Field{
				"email": {
					{
						Name:        []string{"email"},
						MultipleCol: false,
					},
				},

				"id": {
					{
						Name:        []string{"id"},
						MultipleCol: false,
					},
				},
			},

			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRequiredFieldsIndex(tt.headers, tt.requiredFields)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRequiredFieldsIndex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRequiredFieldsIndex() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func createEmployee(data map[string]string, correct bool, reason string) *models.Employee {
	return &models.Employee{
		Data: data,
		Correct: models.CorrectData{
			IsCorrect: correct,
			Reason:    reason,
		},
	}
}
func Test_insertEmpAndCheckUniquiness(t *testing.T) {
	type args struct {
		employee     *models.Employee
		uniqueFields models.UniqueFields
		value        string
		fieldName    string
	}
	tests := []struct {
		name      string
		args      args
		isCorrect bool
	}{
		{
			name: "insert valid employee not inserted yet",
			args: args{
				employee: createEmployee(
					map[string]string{
						"email": "tarcisio_marinho@hotmail.com",
						"name":  "tarcisio",
					},
					true,
					"",
				),
				uniqueFields: models.UniqueFields{
					Fields: map[string]map[string]bool{
						"email": {
							"teste@gmail.com": true,
						},
					},
				},
				value:     "email",
				fieldName: "tarcisio_marinho@hotmail.com",
			},
			isCorrect: true,
		},
		{
			name: "insert valid employee but already inserted",
			args: args{
				employee: createEmployee(
					map[string]string{
						"email": "tarcisio_marinho@hotmail.com",
						"name":  "tarcisio",
					},
					true,
					"",
				),
				uniqueFields: models.UniqueFields{
					Fields: map[string]map[string]bool{
						"email": {
							"tarcisio_marinho@hotmail.com": true,
						},
					},
				},
				value:     "tarcisio_marinho@hotmail.com",
				fieldName: "email",
			},
			isCorrect: false,
		},
		{
			name: "insert empty fieldname",
			args: args{
				employee: createEmployee(
					map[string]string{
						"email": "tarcisio_marinho@hotmail.com",
						"name":  "tarcisio",
					},
					true,
					"",
				),
				uniqueFields: models.UniqueFields{
					Fields: map[string]map[string]bool{
						"email": {
							"ta@fads.com": true,
						},
					},
				},
				value:     "",
				fieldName: "email",
			},
			isCorrect: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertEmpAndCheckUniquiness(tt.args.employee, tt.args.uniqueFields, tt.args.value, tt.args.fieldName)

			if tt.isCorrect {
				if !tt.args.employee.IsCorrect() {
					t.Errorf("insertEmpAndCheckUniquiness() expected = %v, got %v", tt.isCorrect, tt.args.employee.IsCorrect())
					return
				}
			} else {
				if tt.args.employee.IsCorrect() {
					t.Errorf("insertEmpAndCheckUniquiness() expected = %v, got %v", tt.isCorrect, tt.args.employee.IsCorrect())
					return
				}
			}
		})
	}
}
