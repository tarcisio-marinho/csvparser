package validations

import (
	"csvparser/src/models"
	"testing"
)

func getUniqueFieldsMock(values ...string) []string {
	output := make([]string, 0, len(values))
	for _, value := range values {
		output = append(output, value)
	}
	return output
}

func getRequiredFieldsMock(empty int) map[string][]models.Field {
	reqFields := make(map[string][]models.Field)

	if empty == 1 {
		reqFields["email"] = []models.Field{
			{
				Name:        []string{"email"},
				MultipleCol: false,
			},
		}

		reqFields["id"] = []models.Field{
			{
				Name:        []string{"id"},
				MultipleCol: false,
			},
		}

		reqFields["name"] = []models.Field{
			{
				Name:        []string{"f.name", "l.name"},
				MultipleCol: true,
			},
			{
				Name:        []string{"name"},
				MultipleCol: false,
			},
		}
	} else if empty == 2 {
		reqFields := make(map[string][]models.Field)
		reqFields["email"] = []models.Field{
			{
				Name:        []string{},
				MultipleCol: false,
			},
		}

	} else {
		reqFields := make(map[string][]models.Field)
		reqFields["email"] = []models.Field{
			{
				Name:        []string{"", " ", "sadf"},
				MultipleCol: false,
			},
		}
	}
	return reqFields

}

func TestValidateFields(t *testing.T) {
	tests := []struct {
		name    string
		fields  models.Fields
		wantErr bool
	}{
		{
			name: "valid required and unique",
			fields: models.Fields{
				RequiredFields: getRequiredFieldsMock(1),
				UniqueFields:   getUniqueFieldsMock("email", "id"),
			},
			wantErr: false,
		},
		{
			name: "valid required and invalid unique",
			fields: models.Fields{
				RequiredFields: getRequiredFieldsMock(1),
				UniqueFields:   getUniqueFieldsMock("asdf", "id"),
			},
			wantErr: true,
		},
		{
			name: "valid required and unique",
			fields: models.Fields{
				RequiredFields: getRequiredFieldsMock(1),
				UniqueFields:   getUniqueFieldsMock("name", "email", "id"),
			},
			wantErr: false,
		},
		{
			name: "nil required",
			fields: models.Fields{
				RequiredFields: nil,
				UniqueFields:   getUniqueFieldsMock("name", "email", "id"),
			},
			wantErr: true,
		},
		{
			name: "nil unique",
			fields: models.Fields{
				RequiredFields: getRequiredFieldsMock(1),
				UniqueFields:   nil,
			},
			wantErr: true,
		},
		{
			name: "nil required and unique",
			fields: models.Fields{
				RequiredFields: nil,
				UniqueFields:   nil,
			},
			wantErr: true,
		},
		{
			name: "nil field required",
			fields: models.Fields{
				RequiredFields: getRequiredFieldsMock(2),
				UniqueFields:   getUniqueFieldsMock("t1", "t2"),
			},
			wantErr: true,
		},
		{
			name: "empty field required",
			fields: models.Fields{
				RequiredFields: getRequiredFieldsMock(3),
				UniqueFields:   getUniqueFieldsMock("t1", "t2"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateFields(tt.fields); (err != nil) != tt.wantErr {
				t.Errorf("ValidateFields() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
