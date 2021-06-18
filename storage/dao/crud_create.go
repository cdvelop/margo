package dao

import (
	"strings"

	"github.com/cdvelop/margo/storage/config"
	"github.com/cdvelop/margo/storage/share"
	"github.com/cdvelop/margo/storage/transaction"
	"github.com/cdvelop/margo/validate"
)

//CREATE inserta un nuevo registro en la base de datos
func CREATE(object RequestObject, ca CRUDadapter) (mg, dt string, ok bool) {

	if mg, ok = validate.This(object, "new"); !ok {
		return
	}

	db := config.Open(ca)
	tx, cx := transaction.BeginTX(db)

	var sql string
	var arg []interface{}

	sql, mg, arg = MakeSQLINSERT(object, ca)
	// fmt.Printf(">>> SQL INSERT [%v]\n", sql)
	if ok = transaction.ExecuteSQLStatement(sql, arg, tx, cx); !ok {
		tx.Rollback()
		return
	}

	ok = true
	return
}

// MakeSQLINSERT retorna strin sql + id + parametros segun tabla y un maps clave valor
func MakeSQLINSERT(r RequestObject, ca CRUDadapter) (out, id string, arguments []interface{}) { //20.6.27.19:32
	var kname, setValue []string //keys  y $1 $2 o ? ?

	if ido, ok := r.FieldValue()["id_"+r.TableName()]; ok { //agregar id al objeto original si este no existe
		id = ido //id objeto
	} else {
		id = share.GetNewID() //id nuevo
		r.FieldValue()["id_"+r.TableName()] = id
	}
	// log.Printf("def db %v", s.APP.DefaultDatabase)
	var i byte
	for k, v := range r.FieldValue() {
		kname = append(kname, k)
		arguments = append(arguments, v)

		ca.MakeSqInsertSyntax(&i, &setValue)
	}
	// log.Printf("setValue %v", setValue)
	var comas string //comas si son requeridas
	if len(r.FieldValue()) > 1 {
		comas = ", " //comas
	}

	k := strings.Join(kname, comas)
	sv := strings.Join(setValue, comas)
	out = `INSERT INTO ` + r.TableName() + `(` + k + `) VALUES (` + sv + `);`
	return
}
