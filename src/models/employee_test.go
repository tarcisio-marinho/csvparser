package models

import (
	"reflect"
	"testing"
)

func TestCreateEmployee(t *testing.T) {
	tests := []struct {
		name string
		want Employee
	}{
		{"Creating new employee", Employee{Data: make(map[string]string)}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateEmployee(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateEmployee() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_IsCorrect(t *testing.T) {
	type fields struct {
		Data    map[string]string
		Correct CorrectData
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Correct employee",
			fields{
				Data: make(map[string]string),
				Correct: CorrectData{
					IsCorrect: true,
					Reason:    "",
				},
			},
			true,
		},
		{"Incorrect employee",
			fields{
				Data: make(map[string]string),
				Correct: CorrectData{
					IsCorrect: false,
					Reason:    "incorrect employee",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Employee{
				Data:    tt.fields.Data,
				Correct: tt.fields.Correct,
			}
			if got := e.IsCorrect(); got != tt.want {
				t.Errorf("IsCorrect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_SetCorrect(t *testing.T) {
	type fields struct {
		Data    map[string]string
		Correct CorrectData
	}
	tests := []struct {
		name     string
		fields   fields
		expected bool
	}{
		{
			"Set correct if not setted yet",
			fields{
				Data: nil,
				Correct: CorrectData{
					IsCorrect: false,
					Reason:    "",
				},
			},
			true,
		},
		{
			"Set correct if yet setted",
			fields{
				Data: nil,
				Correct: CorrectData{
					IsCorrect: false,
					Reason:    "teste",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			correctValue := tt.fields.Correct.IsCorrect
			e := &Employee{
				Data:    tt.fields.Data,
				Correct: tt.fields.Correct,
			}
			e.SetCorrect()

			if tt.expected != e.IsCorrect() {
				t.Errorf("IsCorrect() = %v, want %v", e.IsCorrect(), correctValue)
			}
		})
	}
}

func TestEmployee_SetIncorrect(t *testing.T) {
	type fields struct {
		Data    map[string]string
		Correct CorrectData
	}
	tests := []struct {
		name            string
		fields          fields
		expectedMessage string
		expectedBool    bool
	}{
		{
			"set incorrect right",
			fields{
				Data:    nil,
				Correct: CorrectData{},
			},
			"message",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Employee{
				Data:    tt.fields.Data,
				Correct: tt.fields.Correct,
			}
			e.SetIncorrect(tt.expectedMessage)

			if tt.expectedBool != e.Correct.IsCorrect {
				t.Errorf("IsCorrect() = %v, want %v", e.Correct.IsCorrect, tt.expectedBool)
			}
		})
	}
}
