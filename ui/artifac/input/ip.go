package input

const patternIp = `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`

// direccion ip valida campos separados por puntos
type Ip struct {
}

//representacion
func (i Ip) Input() (html string) {
	return
}

//validacion con datos de entrada
func (i Ip) Validate(dataIN string) (ok bool) {
	return patternValidate(patternIp, dataIN)
}
