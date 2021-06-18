package input

import "unicode"

const patternOsid = `^([0-9A-Fa-f]{2}[:]){5}([0-9A-Fa-f]{2})$`

// direccion mac valida campos separados por :
type Osid struct {
}

//representacion
func (i Osid) Input() (html string) {
	return patternOsid
}

//validacion con datos de entrada
func (i Osid) Validate(dataIN string) (ok bool) {
	return osidValidate(dataIN)
}

//OSID valida osid formato S-1-5-21-620454579-1741226705-1933409994-1001
//una sola letra
//7 guiones -
//37 numeros
func osidValidate(stringin string) bool {
	var (
		num, let, gui uint8 //numero, letra y guion -
	)

	for _, char := range stringin {
		switch {
		case unicode.IsNumber(char):
			num++
		case unicode.IsLetter(char):
			let++
		case string(char) == "-":
			gui++
		default:
			return false
		}
	}
	if gui != 7 || let != 1 || num != 37 {
		return false
	}
	// log.Printf("nun:%v let:%v guion:%v", num, let, gui)
	return true
}
