package models

import (
	"csvparser/src/utils"
	"strings"
)

type Field struct {
	Name []string // transformar em um objeto ? getName() -> full name ????
	//pra resolver o acesso dessa forma -> info.Name[0]
	// pra quando MultipleCol = true
	MultipleCol bool
}

func (f Field) GetName() string {
	if f.MultipleCol {
		return strings.Join(f.Name, " ")
	}

	if len(f.Name) == 0 {
		return ""
	}

	return f.Name[0]
}

//TODO: e se o campo estiver vazio ? ""
func (f Field) IsValid() bool {
	if len(f.Name) == 0 {
		return false
	}

	if len(f.Name) > 1 && !f.MultipleCol {
		return false
	}

	if f.MultipleCol {
		for _, value := range f.Name {
			if !utils.HasValue(value) {
				return false
			}
		}
	}

	return true
}
