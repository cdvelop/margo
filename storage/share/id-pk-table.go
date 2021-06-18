package share

import (
	"strings"
)

//IdpkTABLA recibe kname y tabla name retorna 2 true si: es fk  y  pkey de la tabla
func IdpkTABLA(keyNameIN, tableName string) (primarykey, primaryKeyThisTable bool) {
	keyname := strings.ToLower(keyNameIN)
	if len(keyname) >= 2 {

		if keyname[:2] != "id" {
			return
		}

		primarykey = true
		keymenosguion := strings.Replace(keyname, "_", "", -1) //remover _
		keynemosid := keymenosguion[2:]                        //remover id

		if keynemosid == tableName {
			primaryKeyThisTable = true
		}
	}
	return
}
