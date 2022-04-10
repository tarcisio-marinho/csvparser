package models

type Field struct {
	Name []string // transformar em um objeto ? getName() -> full name ????
	//pra resolver o acesso dessa forma -> info.Name[0]
	// pra quando MultipleCol = true
	MultipleCol bool
}
