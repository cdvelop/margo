package input

const patternTextNum = `^[A-Za-z0-9_]{5,15}$`

// texto, numero y guion bajo 5 a 15 caracteres
type TextNum struct {
}

//representacion
func (r TextNum) Input() (html string) {
	return
}

//validacion con datos de entrada
func (r TextNum) Validate(dataIN string) (ok bool) {
	return patternValidate(patternTextNum, dataIN)
}
