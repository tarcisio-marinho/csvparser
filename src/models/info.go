package models

//TODO: validar se essa é a melhor forma de armazenar esse dado
type Info struct {
	Name []string // transformar em um objeto ? getName() -> full name ????
	//pra resolver o acesso dessa forma -> info.Name[0]
	// pra quando MultipleCol = true
	MultipleCol bool
	Index       int
}

//TODO: criar alguma forma de adicionar mais campos aqui
//TODO: arrumar algum jeito de pegar múltiplas colunas como sendo um único campo ->
// será que posso separar por vírgula a string ? "primeiro_campo segundo_campo terceiro_campo" ?
func CreateFields() map[string][]Info {
	fields := make(map[string][]Info)
	fields["name"] = []Info{
		{Name: []string{"name"}, MultipleCol: false},
		{Name: []string{"f.name", "l.name"}, MultipleCol: true},
		{Name: []string{"firstname", "lastname"}, MultipleCol: true},
		{Name: []string{"first", "last"}, MultipleCol: true},
	}

	fields["email"] = []Info{
		{Name: []string{"email"}, MultipleCol: false},
		{Name: []string{"e-mail"}, MultipleCol: false},
	}

	fields["id"] = []Info{
		{Name: []string{"id"}, MultipleCol: false},
		{Name: []string{"number"}, MultipleCol: false},
		{Name: []string{"employeenumber"}, MultipleCol: false},
		{Name: []string{"empid"}, MultipleCol: false},
	}

	fields["salary"] = []Info{
		{Name: []string{"salary"}, MultipleCol: false},
		{Name: []string{"wage"}, MultipleCol: false},
		{Name: []string{"rate"}, MultipleCol: false},
	}

	return fields
}
