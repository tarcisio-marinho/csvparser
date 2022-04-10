package models

//TODO: colocar employee em um pacote separado ?
type Employee struct {
	Data    map[string]string
	Correct bool
}

func CreateEmployee() Employee {
	return Employee{Data: make(map[string]string)}
}

func (e Employee) IsCorrect() bool {
	return e.Correct
}

/* DEPRECATED
func (e Employee) IsCorrect() bool {
	for _, data := range e.Data {
		if !utils.HasValue(data) {
			return false
		}
	}

	return true
}
*/
