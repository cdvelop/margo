package input

type Radio struct {
	Values []string //{"Dama","Varon"}
}

//NewRadio "Varon", "Dama"
func NewRadio(valuesIN ...string) Radio {
	newR := Radio{
		Values: valuesIN,
	}
	return newR
}

//representacion en html  "Dama" = {"D:Dama"} ; "Varon" = {"V":"Varon"}
func (r Radio) Input() (html string) {
	return
}

//validacion con datos de entrada
func (r Radio) Validate(dataIN string) (ok bool) {

	if len(dataIN) <= 0 {
		return
	}

	if len(dataIN) > 1 {
		return
	}

	for _, value := range r.Values {
		if value[0:1] == dataIN {
			ok = true
			return
		}
	}
	return
}
