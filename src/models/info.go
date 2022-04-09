package models

import "csvparser/src/utils"

//TODO: colocar employee em um pacote separado ?
type Employee struct {
	Data map[string]string
}

func CreateEmployee() Employee {
	return Employee{Data: make(map[string]string)}
}

func (e Employee) IsCorrect() bool {
	for _, data := range e.Data {
		if !utils.HasValue(data) {
			return false
		}
	}

	return true
}

type FieldIndex struct {
	FieldName   string
	Index       []int
	MultipleCol bool
}

// TODO: é essa a melhor estrutura ?
type RequiredFields struct {
	FieldName string
	Fields    map[string][]Info
}

func CreateRequiredFields() RequiredFields {
	return RequiredFields{Fields: make(map[string][]Info)}
}

//TODO: validar se essa é a melhor forma de armazenar esse dado
type Info struct {
	Name []string // transformar em um objeto ? getName() -> full name ????
	//pra resolver o acesso dessa forma -> info.Name[0]
	// pra quando MultipleCol = true
	MultipleCol bool
}

//TODO: criar alguma forma de adicionar mais campos aqui
//TODO: arrumar algum jeito de pegar múltiplas colunas como sendo um único campo ->
// será que posso separar por vírgula a string ? "primeiro_campo segundo_campo terceiro_campo" ?
func CreateFields() RequiredFields {
	reqFields := CreateRequiredFields()

	reqFields.Fields["name"] = []Info{
		{Name: []string{"name"}, MultipleCol: false},
		{Name: []string{"f.name", "l.name"}, MultipleCol: true},
		{Name: []string{"firstname", "lastname"}, MultipleCol: true},
		{Name: []string{"first", "last"}, MultipleCol: true},
	}

	reqFields.Fields["email"] = []Info{
		{Name: []string{"email"}, MultipleCol: false},
		{Name: []string{"e-mail"}, MultipleCol: false},
	}

	reqFields.Fields["id"] = []Info{
		{Name: []string{"id"}, MultipleCol: false},
		{Name: []string{"number"}, MultipleCol: false},
		{Name: []string{"employeenumber"}, MultipleCol: false},
		{Name: []string{"empid"}, MultipleCol: false},
	}

	reqFields.Fields["salary"] = []Info{
		{Name: []string{"salary"}, MultipleCol: false},
		{Name: []string{"wage"}, MultipleCol: false},
		{Name: []string{"rate"}, MultipleCol: false},
	}

	return reqFields
}

/* TODO: implement
func AddField(requiredFieldName string, columns []string, isMultipleColumn bool) {
	i[requiredFieldName] = Info{
		Name:        columns,
		MultipleCol: isMultipleColumn,
	}
}
*/
