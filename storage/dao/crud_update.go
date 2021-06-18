package dao

import (
	"strings"

	"github.com/cdvelop/margo/storage/config"
	"github.com/cdvelop/margo/storage/share"
	"github.com/cdvelop/margo/storage/transaction"
	"github.com/cdvelop/margo/validate"
)

//UPDATE inserta o actualiza un obj en la base de datos
func UPDATE(object RequestObject, ca CRUDadapter) (mg, dt string, ok bool) {

	if mg, ok = validate.This(object, ""); !ok {
		return
	}

	db := config.Open(ca)
	tx, cx := transaction.BeginTX(db)

	var sql string
	var arg []interface{}

	sql, mg, arg = MakeSQLUPDATE(object, ca)

	if ok = transaction.ExecuteSQLStatement(sql, arg, tx, cx); !ok {
		tx.Rollback()
		return
	}
	ok = true
	return
}

//MakeSQLUPDATE genera un strin sql segun tabla y un maps clave valor
func MakeSQLUPDATE(ro RequestObject, ca CRUDadapter) (sql, idval string, values []interface{}) { //20.6.29.19:32
	var (
		set     []string //set nombre=$2, etc
		c, idpk string
	)
	var i byte
	for k, v := range ro.FieldValue() {
		if _, pk := share.IdpkTABLA(k, ro.TableName()); !pk {
			i++

			ca.SetListSyntax(k, i, &set)

			values = append(values, v)
		} else {
			idpk = k
			idval = v
		}
	}

	//ultimo argumento valor id pk del objeto
	values = append(values, idval)
	if len(ro.FieldValue()) > 1 {
		c = ", " //comas
	}
	st := strings.Join(set, c)
	sql = `UPDATE ` + ro.TableName() + ` SET ` + st + ` WHERE ` + idpk + ` = ` + ca.TotalValuesSyntax(ro.FieldValue()) + `;`
	//sql = `UPDATE jugador SET idjugador=$1, nombre=$2, apellido=$3, genero=$4 WHERE idjugador = $5;`
	// log.Printf("sql update: %v values %v", sql, values)
	// log.Printf("argumentos: %v", values)
	return
}
