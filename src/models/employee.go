package models

import "csvparser/src/utils"

type Employee struct {
	Data    map[string]string
	Correct CorrectData
}

type CorrectData struct {
	IsCorrect bool
	Reason    string
}

func (e *Employee) SetIncorrect(message string) {
	e.Correct.IsCorrect = false
	e.Correct.Reason = message
}

func (e *Employee) SetCorrect() {
	if utils.HasValue(e.Correct.Reason) {
		return
	} else {
		e.Correct.IsCorrect = true
	}
}

func CreateEmployee() Employee {
	return Employee{Data: make(map[string]string)}
}

func (e Employee) IsCorrect() bool {
	return e.Correct.IsCorrect
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
