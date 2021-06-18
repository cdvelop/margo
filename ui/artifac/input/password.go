package input

const patternPassword = `[a-zA-Z][a-zA-Z0-9-_.]{5,10}`

// [a-zA-Z][a-zA-Z0-9-_.]{5,15}
//  Solo letras (en cualquier caso), números, guiones, guiones bajos y puntos.
//  (No el carácter de barra, que se usa para escapar del punto).
// debe comenzar con una letra y debe tener entre 5 y 10 caracteres (inclusive).
type Password struct {
}

//representacion
func (p Password) Input() (html string) {
	return "contraseña minimo 8 caracteres"
}

//validacion con datos de entrada
func (p Password) Validate(dataIN string) (ok bool) {
	return patternValidate(patternPassword, dataIN)
}
