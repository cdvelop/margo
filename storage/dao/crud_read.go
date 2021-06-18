package dao

import (
	"fmt"
	"strings"

	"github.com/cdvelop/margo/storage/share"
)

//READallColumnsByID
func READallColumnsByID(da DbAdapter, ro RequestObject) (out map[string]string, ok bool) {
	sql := MakeSQLSELECT("*", ro.TableName(), ro.FieldValue())
	if out, _ = QueryOne(sql, da); len(out) != 0 {
		ok = true
	} else {
		ok = false
	}
	return
}

// MakeSQLSELECT s(count(*),) tabla(jugador,usuario) id(2,7,12) genera un strin sql segun tabla y un maps clave valor
func MakeSQLSELECT(option, TNAME string, obj map[string]string) (out string) { //20.6.17.19:32
	var (
		selec, que        []string
		js, jw, pk, where string
	)
	for columnName, value := range obj {
		if value != "" {
			where = ` WHERE `
			if _, pkt := share.IdpkTABLA(columnName, TNAME); !pkt {
				selec = append(selec, columnName)
				que = append(que, fmt.Sprintf("%v ='%v'", columnName, value))
			} else { //hay pk en el objeto
				pk = value
			}
		} else {
			selec = append(selec, columnName)
		}
	}

	if len(obj) > 1 {
		js = ", "
		jw = " AND "
	}

	if len(option) == 0 { //si el select trae
		option = strings.Join(selec, js)

	}

	if pk != "" { //si no esta vacio pk
		que = nil //anulo la data recogida
		que = append(que, fmt.Sprintf("id_%v ='%v'", TNAME, pk))
	}

	q := strings.Join(que, jw)
	out = fmt.Sprintf("SELECT %v FROM %v%v%v;", option, TNAME, where, q)
	// fmt.Printf(">>> SQL SELECT: [%v]\n", out)
	return
}
