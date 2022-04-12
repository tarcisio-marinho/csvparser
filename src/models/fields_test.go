package models

import (
	"reflect"
	"testing"
)

func GetRequiredFieldsMock() map[string][]Field {
	reqFields := make(map[string][]Field)
	reqFields["email"] = []Field{
		{
			Name:        []string{"email"},
			MultipleCol: false,
		},
	}

	reqFields["id"] = []Field{
		{
			Name:        []string{"id"},
			MultipleCol: false,
		},
	}

	reqFields["name"] = []Field{
		{
			Name:        []string{"f.name", "l.name"},
			MultipleCol: true,
		},
		{
			Name:        []string{"name"},
			MultipleCol: false,
		},
	}
	return reqFields
}

func TestCreateFieldsFromConfig(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		want    Fields
		wantErr bool
	}{
		{
			name:    "empty data, empty byte array",
			bytes:   nil,
			want:    Fields{},
			wantErr: true,
		},
		{
			name:    "empty data, empty byte array",
			bytes:   []byte{},
			want:    Fields{},
			wantErr: true,
		},
		{
			name: "correct fields, required and unique",
			bytes: []byte(
				`{
    "RequiredFields": {
        "email": [
        {
            "Name": [
                "email"
            ],
            "MultipleCol": false
        }       
        ],
        "id": [
        {
            "Name": [
                "id"
            ],
            "MultipleCol": false
        }],
        "name": [
        {
            "Name": [
                "f.name",
                "l.name"
            ],
            "MultipleCol": true
        },
        {
            "Name": [
                "name"
            ],
            "MultipleCol": false
        }]
    },
    "UniqueFields": [
            "email",
            "id"
        ]
    }`),
			want: Fields{
				RequiredFields: GetRequiredFieldsMock(),
				UniqueFields:   []string{"email", "id"},
			},
			wantErr: false,
		},
		{
			name: "incorrect json data",
			bytes: []byte(
				`{
    "RequiredFields": {
        "email": [
        {
            "Name": [],
            "MultipleCol": false
        }       ultipleCol": false
        }],
        "name": [
        {
            "Name": [],
            "MultipleCol": true
        },
        {
            "Name": [],
            "MultipleCol": false
        }]
    },
    "UniqueFields": [
            "email",
            "id"
        ]
    }`),
			want: Fields{
				RequiredFields: nil,
				UniqueFields:   nil,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateFieldsFromConfig(tt.bytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateFieldsFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateFieldsFromConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
