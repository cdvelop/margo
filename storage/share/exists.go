package share

import "fmt"

//Exists verifica si existe un que es y algo en la base de datos segun sql
//ej: "la base de datos","tiendadb","sql para la consulta"
func Exists(textResponse, objectSelect, sql string, d DbAdapter) (ok bool) {
	var val string

	if val, ok = SelectValue(sql, d); ok && val == "1" {
		fmt.Printf(">>> ok %v %v existe\n", textResponse, objectSelect)
	} else {
		fmt.Printf("!!! error %v %v no existe\n", textResponse, objectSelect)
	}
	return
}
