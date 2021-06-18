package input

// solo texto. minimo 3 caracteres no numeros
type Text struct {
}

//representacion
func (r Text) Input() (html string) {
	return `solo texto. minimo 3 caracteres no numeros`
}

//validacion con datos de entrada
func (r Text) Validate(dataIN string) (ok bool) {
	pattern := `^[^0-9 ]{1}([A-Za-zÑñ ])+[A-Za-zÑñ]+$`
	return patternValidate(pattern, dataIN)
}
