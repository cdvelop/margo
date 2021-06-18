package input

const patternNumber = `^[0-9]\d*$`

// solo valores numericos positivos >= 0
type Number struct {
}

//representacion
func (r Number) Input() (html string) {
	return `solo valores numericos positivos >= 0`
}

//validacion con datos de entrada
func (r Number) Validate(dataIN string) (ok bool) {
	return patternValidate(patternNumber, dataIN)
}
