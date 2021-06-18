package migrate

import (
	"fmt"
	"log"

	"github.com/cdvelop/margo/storage/config"
)

//DeleteDataFromTABLE borra data de una tabla en db
func DeleteDataFromTABLE(tableName string, dba DbAdapter) {
	db := config.Open(dba)
	sql := fmt.Sprintf("DELETE FROM %v;", tableName)
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
}

//DeleteTABLE elimina tabla de una base de datos
func DeleteTABLE(tableName string, dba DbAdapter) {
	db := config.Open(dba)
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %v;", tableName)
	if _, err := db.Exec(sql); err != nil {
		log.Fatal(err)
	}
}
