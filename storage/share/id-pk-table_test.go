package share_test

import (
	"testing"

	"github.com/cdvelop/margo/storage/share"
)

func Test_IdpkTABLA(t *testing.T) {
	// tableName := "user"

	datatest := map[string]struct {
		tableName           string
		keyNameIN           string
		primarykey          bool
		primaryKeyThisTable bool
	}{
		"id corresponde a tabla sin guion":                               {tableName: "usuario", keyNameIN: "idusuario", primarykey: true, primaryKeyThisTable: true},
		"id corresponde a tabla guion bajo _":                            {tableName: "especialidad", keyNameIN: "id_especialidad", primarykey: true, primaryKeyThisTable: true},
		"corresponde a tabla y key contiene parte el nombre de la tabla": {tableName: "especialidad", keyNameIN: "especialidades", primarykey: false, primaryKeyThisTable: false},
		"id fk de otra tabla sin guion":                                  {tableName: "usuario", keyNameIN: "idfactura", primarykey: true, primaryKeyThisTable: false},
		"id fk de otra tabla con guion bajo _":                           {tableName: "usuario", keyNameIN: "id_factura", primarykey: true, primaryKeyThisTable: false},
		"no primary key presente":                                        {tableName: "usuario", keyNameIN: "factura", primarykey: false, primaryKeyThisTable: false},
		"nenos de 2 caracteres id no presente":                           {tableName: "usuario", keyNameIN: "i", primarykey: false, primaryKeyThisTable: false},
	}

	for testName, dt := range datatest {

		t.Run((testName), func(t *testing.T) {
			pk, pkthistable := share.IdpkTABLA(dt.keyNameIN, dt.tableName)

			if pk != dt.primarykey {
				t.Fail()
			}
			if pkthistable != dt.primaryKeyThisTable {
				t.Fail()
			}
		})

	}
}
