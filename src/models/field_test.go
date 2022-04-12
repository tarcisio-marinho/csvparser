package models

import "testing"

func TestField_GetName(t *testing.T) {
	type fields struct {
		Name        []string
		MultipleCol bool
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"testing single column valid name",
			fields{
				Name:        []string{"tarcisio"},
				MultipleCol: false,
			},
			"tarcisio",
		},
		{"testing single column empty name",
			fields{
				Name:        []string{""},
				MultipleCol: false,
			},
			"",
		},
		{"testing single column empty array name",
			fields{
				Name:        nil,
				MultipleCol: false,
			},
			"",
		},
		{"testing multiple column name",
			fields{
				Name:        []string{"tarcisio", "marinho"},
				MultipleCol: true,
			},
			"tarcisio marinho",
		},
		{"testing more than two column name",
			fields{
				Name:        []string{"tarcisio", "marinho", "de", "oliveira"},
				MultipleCol: true,
			},
			"tarcisio marinho de oliveira",
		},
		{"testing multiple column empty array",
			fields{
				Name:        nil,
				MultipleCol: true,
			},
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Field{
				Name:        tt.fields.Name,
				MultipleCol: tt.fields.MultipleCol,
			}
			if got := f.GetName(); got != tt.want {
				t.Errorf("GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestField_IsValid(t *testing.T) {
	type fields struct {
		Name        []string
		MultipleCol bool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "valid multiple column name",
			fields: fields{
				Name:        []string{"tarcisio", "marinho"},
				MultipleCol: true,
			},
			want: true,
		},
		{
			name: "valid single column name",
			fields: fields{
				Name:        []string{"tarcisio"},
				MultipleCol: false,
			},
			want: true,
		},
		{
			name: "valid with more than two column name",
			fields: fields{
				Name:        []string{"tarcisio", "marinho", "oliveira"},
				MultipleCol: true,
			},
			want: true,
		},
		{
			name: "invalid without fields",
			fields: fields{
				Name:        nil,
				MultipleCol: true,
			},
			want: false,
		},
		{
			name: "invalid with empty fields",
			fields: fields{
				Name:        []string{},
				MultipleCol: true,
			},
			want: false,
		},
		{
			name: "invalid with empty fields",
			fields: fields{
				Name:        []string{},
				MultipleCol: true,
			},
			want: false,
		},
		{
			name: "invalid with multiple column and multcol: false",
			fields: fields{
				Name:        []string{"tarcisio", "marinho"},
				MultipleCol: false,
			},
			want: false,
		},
		{
			name: "invalid with empty multiple column",
			fields: fields{
				Name:        []string{"", ""},
				MultipleCol: false,
			},
			want: false,
		},
		{
			name: "invalid with only spaces multiple column",
			fields: fields{
				Name:        []string{" ", " "},
				MultipleCol: false,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Field{
				Name:        tt.fields.Name,
				MultipleCol: tt.fields.MultipleCol,
			}
			if got := f.IsValid(); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
