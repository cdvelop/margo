package migrate

import (
	"fmt"
	"log"
	"strings"

	"github.com/cdvelop/margo/storage/config"
	"github.com/cdvelop/margo/storage/transaction"
)

var err error

//ClonDATATABLE copia la data de una tabla a otra nueva
func ClonDATATABLE(tableName string, fields []map[string]string, da DbAdapter, o ORM) (out bool) {
	var toClone []string //columnas a copiar
	var ok bool
	tmp := `tabtemp`

	db := config.Open(da) //obtener conexion
	// `tx` es una instancia de` * sql.Tx` a través de la cual podemos ejecutar nuestras consultas
	tx, ctx := transaction.BeginTX(db) // Crea un nuevo contexto y comienza una transacción
	defer tx.Rollback()

	// 1 renombrar tabla
	_, err = tx.ExecContext(ctx, `ALTER TABLE `+tableName+` RENAME TO `+tmp+`;`)
	if err != nil {
		log.Printf("!!! error %v al renonbrar tabla %v", err, tableName)
		tx.Rollback()
		return
	}
	//2 crear tabla nueva
	sql := makeSQLCreaTABLE(tableName, fields)
	_, err = tx.ExecContext(ctx, sql)
	if err != nil {
		tx.Rollback()
		return
	}
	//3 seleccionar data anterior
	var oldfield []string
	q := fmt.Sprintf(o.SQLTableInfo(), tableName)

	// knames, ok = tx.getallOBJ(&q, &ctx)
	var knames = make([]map[string]string, 0)
	if knames, ok = transaction.SelectAllContext(&q, db.DB, &ctx); !ok { //entrega nombre columnas de la tabla
		tx.Rollback()
		return
	}
	log.Printf(">>> columnas selecionadas para copiar: %v", knames)

	for _, d := range knames {
		oldfield = append(oldfield, d[o.SQLColumName()])
	}

	for _, field := range fields {
		for fieldName := range field {

			for _, oldfield := range oldfield {
				if fieldName == oldfield {
					toClone = append(toClone, oldfield) //agrego las columnas que seran copiadas
					break
				}
			}

		}
	}

	log.Printf(">>> toClone %v", toClone)
	//4 copiar data
	c := strings.Join(toClone, ",") //creando un string separado por ,
	// INSERT INTO ciudad (idciudad,nombre) SELECT idciudad,nombre FROM temp
	q = fmt.Sprintf("INSERT INTO %v (%v) SELECT %v FROM %v;", tableName, c, c, tableName)
	log.Printf(">>> sql insert %v", q)
	log.Printf(">>> copiando data %v", tableName)
	_, err = tx.ExecContext(ctx, q)
	if err != nil {
		log.Printf("!!! error %v al copiar data de %v a tabla %v", err, tmp, tableName)
		tx.Rollback()
		return
	}
	log.Printf(">>> data copiada: %v", tableName)

	//5 borrar tabla temporal
	q = fmt.Sprintf(o.SQLDropTable(), tableName)
	// log.Printf(">> sql droptab : %v", q)
	_, err = tx.ExecContext(ctx, q)
	if err != nil {
		log.Printf("!!! error %v al borrar tabla temporal %v", err, tableName)
		tx.Rollback()
		return
	}

	// Finalmente, si no se reciben errores de las consultas, confirme la transacción
	// esto aplica los cambios anteriores a nuestra base de datos
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Fatal(err)
		return
	}
	out = true
	return
}
