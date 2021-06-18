package input

import (
	"fmt"
	"strconv"
	"strings"
)

const patternRut = `^[0-9]+-[0-9kK]{1}$`

// rut sin puntos con guion ejem.: 11222333-4
type Rut struct {
}

//representacion
func (r Rut) Input() (html string) {
	fmt.Println(patternRut)
	return `formato "7863697-1`
}

//validacion con datos de entrada
func (r Rut) Validate(dataIN string) bool {
	return runValidate(dataIN)
}

//RUT validate formato "7863697-1"
func runValidate(rin string) bool {
	r := strings.Split(string(rin), "-")
	var mint int

	var err error
	if mint, err = strconv.Atoi(r[0]); err != nil {
		return false
	}

	dv := dvrut(mint)

	return dv == r[1]
}

// dvrut retorna digito verificador de un run
func dvrut(rut int) string {
	var sum = 0
	var factor = 2
	for ; rut != 0; rut /= 10 {
		sum += rut % 10 * factor
		if factor == 7 {
			factor = 2
		} else {
			factor++
		}
	}

	if val := 11 - (sum % 11); val == 11 {
		return "0"
	} else if val == 10 {
		return "k"
	} else {
		return strconv.Itoa(val)
	}
}
