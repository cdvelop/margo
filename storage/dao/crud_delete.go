package dao

import (
	"fmt"

	"github.com/cdvelop/margo/storage/config"
	"github.com/cdvelop/margo/storage/share"
	"github.com/cdvelop/margo/storage/transaction"
)

//DELETE borra objeto en base de datos..
func DELETE(object RequestObject, ca CRUDadapter) (mg, dt string, ok bool) {
	sql, arg := MakeSQLDELETE(object, ca)

	// sql, arg := MakeSQLDELETE(object, da, orm)
	// log.Printf("sql delete: %v arg: %v", sql, arg)

	db := config.Open(ca)
	tx, cx := transaction.BeginTX(db)
	defer tx.Rollback()

	if ok = transaction.ExecuteSQLStatement(sql, arg, tx, cx); !ok {
		tx.Rollback()
		return
	}
	mg = object.TableName() + " eliminad@ correctamente"
	dt = fmt.Sprintf(`{"id":"%v"}`, arg)
	ok = true
	return
}

// MakeSQLDELETE retonna sql string para eliminacion
// func MakeSQLDELETE(object RequestObject, da DbAdapter, orm ORM) (out string, id []interface{}) {
func MakeSQLDELETE(object RequestObject, ca CRUDadapter) (out string, id []interface{}) {

	for k, v := range object.FieldValue() {
		if _, pk := share.IdpkTABLA(k, object.TableName()); pk {
			out = fmt.Sprintf("DELETE FROM %v WHERE %v = %v", object.TableName(), k, ca.SQLsyntax())
			id = append(id, v)
			break
		}
	}
	return
}
