package share

import (
	"fmt"
	"strings"
)

//ObjToJSON retorna un string con formato json de un objeto quitando pk
func ObjToJSON(tname string, obj *map[string]string) (mg, dt string) {
	var (
		pke  int      //primary key encontrado
		c    string   //separacion con comas si es necesario
		m, d []string //array de mjes y data complementaria
	)

	for k, v := range *obj {
		if _, pk := IdpkTABLA(k, tname); !pk { //no pk de la tabla
			d = append(d, fmt.Sprintf(`{"name":"%v","value": "%v"}`, k, v))
			m = append(m, fmt.Sprintf(" %v %v ", k, v))
		} else { //objeto tiene pk
			pke = 1
		}
	}

	if len(*obj)-pke > 1 {
		c = ", " //separacion array
	}

	dt = "[" + strings.Join(d, c) + "]"
	mg = strings.Join(m, c)
	return
}
