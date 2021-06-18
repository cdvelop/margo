package dao

import (
	"fmt"
	"strings"
)

//GetConcatenateTITLES retorna name y value concatenado un objeto por idname en caso de tablas foreaneas
//titles ej: ["nombre","apellido"] campos que se mostraran al usuario en un listado
func GetConcatenateTITLES(idname string, titles []string, da DbAdapter) (lst []map[string]string) {
	if len(idname) > 1 {
		var (
			c string //string de concatenacion
		)
		tabla := (idname)[2:] //tabla foranea restando palabra id

		if len(titles) > 1 {
			c = ` || ' ' || `
		}
		s := strings.Join(titles, c)
		// log.Printf("tabla %v array %v tamaño %v", tabla, tl, len(tl))
		sql := makeSQLSELNAMEVALUE(s, idname, tabla)
		// log.Printf("sql: %v", sql)
		lst, _ = QueryAll(sql, da)
	}
	return
}

//makeSQLSELNAMEVALUE retorna sql string = SELECT nombre || ' ' || apellido AS name, idjugador AS id FROM jugador;
func makeSQLSELNAMEVALUE(s, idname, tname string) (out string) {
	out = fmt.Sprintf("SELECT %v AS name, %v AS value FROM %v;", s, idname, tname)
	return
}
