package models

// TODO: é essa a melhor estrutura ?
type RequiredFields struct {
	Fields map[string][]Field
}

func CreateRequiredFields() RequiredFields {
	return RequiredFields{Fields: make(map[string][]Field)}
}

// DEPRECATED
func CreateFields() RequiredFields {
	reqFields := CreateRequiredFields()

	reqFields.Fields["name"] = []Field{
		{Name: []string{"name"}, MultipleCol: false},
		{Name: []string{"f.name", "l.name"}, MultipleCol: true},
		{Name: []string{"firstname", "lastname"}, MultipleCol: true},
		{Name: []string{"first", "last"}, MultipleCol: true},
	}

	reqFields.Fields["email"] = []Field{
		{Name: []string{"email"}, MultipleCol: false},
		{Name: []string{"e-mail"}, MultipleCol: false},
	}

	reqFields.Fields["id"] = []Field{
		{Name: []string{"id"}, MultipleCol: false},
		{Name: []string{"number"}, MultipleCol: false},
		{Name: []string{"employeenumber"}, MultipleCol: false},
		{Name: []string{"empid"}, MultipleCol: false},
	}

	reqFields.Fields["salary"] = []Field{
		{Name: []string{"salary"}, MultipleCol: false},
		{Name: []string{"wage"}, MultipleCol: false},
		{Name: []string{"rate"}, MultipleCol: false},
	}

	return reqFields
}
