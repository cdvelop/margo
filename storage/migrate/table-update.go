package migrate

import (
	"fmt"
	"log"

	"github.com/cdvelop/margo/storage/dao"
)

//UpdateTABLES revisa si tienen data las tablas para actualizarlas y respaldar la data
func UpdateTABLES(da DbAdapter, o ORM, Tables ...map[string][]map[string]string) bool {
	for _, table := range Tables {
		for tableName, fields := range table {

			//consulta entrega columna nombre
			q := fmt.Sprintf(o.SQLTableInfo(), tableName)
			list, ok := dao.QueryOne(q, da)
			if !ok {
				return false
			}
			// ok = Err(err)

			if len(list) == 0 { //si no existe crear tabla nueva
				CreateOneTABLE(da, tableName, fields)
			} else { //revisar tabla consultar si tiene data

				if list, ok := dao.QueryOne("SELECT * FROM "+tableName+";", da); ok {

					if len(list) == 0 { //lista sin data borramos tabla y la creamos nuevamente para no chequearla
						q := fmt.Sprintf(o.SQLDropTable(), tableName)
						if dao.QueryWithoutANSWER(q, ">>> borrando tabla "+tableName+"", da) {
							log.Printf(">>> tabla %v sin data borrada", tableName)
							if !CreateOneTABLE(da, tableName, fields) {
								return false
							}
							log.Printf(">>> tabla %v creada", tableName)
						} else {
							log.Printf("!!! error al borrar tabla DROP TABLE: " + tableName)
							return false
						}

					} else { //lista con data hay que actualizar
						// log.Printf("tabla %v con data. hay que verificar", tableName)
						if !ClonDATATABLE(tableName, fields, da, o) { //clonamos la tabla con data a la nueva
							log.Printf("!!! error al copiar la data tabla " + tableName)
							return false
						}
					}

				} else {
					return false
				}
			}
		}
	} //*****tablas*****

	return true
}
