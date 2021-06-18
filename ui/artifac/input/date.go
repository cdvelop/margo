package input

const patternDate = `[0-9]{4}-(0[1-9]|1[012])-(0[1-9]|1[0-9]|2[0-9]|3[01])`

// formato fecha: DD-MM-YYYY
type Date struct {
}

//representacion
func (r Date) Input() (html string) {
	return `formato fecha: DD-MM-YYYY`
}

//validacion con datos de entrada
func (r Date) Validate(dataIN string) (ok bool) { //en realidad es YYYY-MM-DD
	if len(dataIN) > 10 {
		return
	}

	return patternValidate(patternDate, dataIN)
}
