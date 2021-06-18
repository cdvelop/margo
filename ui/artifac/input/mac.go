package input

const patternMac = `^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$`

// direccion mac valida campos separados por :
type Mac struct {
}

//representacion
func (i Mac) Input() (html string) {
	return
}

//validacion con datos de entrada
func (i Mac) Validate(dataIN string) (ok bool) {
	return patternValidate(patternMac, dataIN)
}
