package migrate

import (
	"fmt"
	"log"
	"strings"

	"github.com/cdvelop/margo/storage/dao"
)

//CreateAllTABLES crea todas las tablas de la base de datos
func CreateAllTABLES(da DbAdapter, tables ...Table) (ok bool) {
	var sql []string
	//todo el sql por tabla
	for _, table := range tables { //array tablas ordenadas
		// fmt.Printf("Nombre de la tabla [%v]\n", table.Name())
		sql = append(sql, makeSQLCreaTABLE(table.Name(), table.FieldsDataType()))
	}

	q := strings.Join(sql, "\n")
	// log.Printf(">>> sql final %v", q)
	if !dao.QueryWithoutANSWER(q, ">>> creando tablas modelo en db", da) {
		return
	}

	log.Printf(">>> tablas db creadas")
	ok = true
	return
}

//CreateOneTABLE segun nombre tabla y solo con un id_nombretabla correlativo por defecto
func CreateOneTABLE(da DbAdapter, tableName string, fields []map[string]string) bool {
	sql := makeSQLCreaTABLE(tableName, fields)
	return dao.QueryWithoutANSWER(sql, ">>> creando tabla: "+tableName+" en db", da)
}

//makeSQLCreaTABLE crea string sql crea tabla
func makeSQLCreaTABLE(tableName string, fields []map[string]string) (sql string) {
	keyLisTO, _ := createSqlListByField(tableName, fields...)

	column := strings.Join(keyLisTO, ", ")
	// fmt.Printf("colum tabla: %v  %v\n", tableName, column)
	return makesqlcretetable(tableName, column)
}

//sql de creacion de tabla
func makesqlcretetable(tname, column string) (sql string) {
	sql = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v (%v);", tname, column)
	return
}
