package input

import (
	"strings"
)

//Check Box id y valor
type Check struct {
	Values map[string]string //{"1":"Admin","2":"Editor"}
}

//representacion en html
func (r Check) Input() (html string) {
	return
}

//validacion con datos de entrada
func (r Check) Validate(dataIN string) (ok bool) {
	dataInArray := strings.Split(dataIN, ",")
	var keysFound []string
	var badKey []string
	for _, idkeyIn := range dataInArray {
		if _, exists := r.Values[idkeyIn]; exists {
			keysFound = append(keysFound, idkeyIn)
		} else {
			badKey = append(badKey, idkeyIn)
		}
	}
	if len(keysFound) > 0 && len(badKey) == 0 {
		ok = true
	}
	return
}
