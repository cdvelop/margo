package input

const patternHour = `^([0-9]|0[0-9]|1[0-9]|2[0-3]):[0-5][0-9]$`

// formato fecha: DD-MM-YYYY
type Hour struct {
}

//representacion
func (r Hour) Input() (html string) {
	return `formato hora: HH:MM`
}

//validacion con datos de entrada
func (r Hour) Validate(dataIN string) (ok bool) { //en realidad es YYYY-MM-DD
	if len(dataIN) > 10 {
		return
	}

	return patternValidate(patternHour, dataIN)
}
