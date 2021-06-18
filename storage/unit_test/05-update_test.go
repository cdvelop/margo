package unit_test

import (
	"testing"

	"github.com/cdvelop/margo/model"
	"github.com/cdvelop/margo/storage/dao"
)

func Update(db dao.CRUDadapter, t *testing.T) {

	for _, dataRead := range dataTestCRUD {

		if dataRead.result { //solo los casos de exito
			updateObject := model.Object{
				Table_Name: defaulTableName,
				Field_Value: map[string]string{
					"id_usuario": dataRead.idRecovered,
					"apellido":   "Actualizado de las carmenes"},
			}

			t.Run(("UPDATE: " + dataRead.idRecovered), func(t *testing.T) {

				mg, dt, ok := dao.UPDATE(updateObject, db)
				if !ok {
					t.Logf("message: %v data %v ok[%v]\n", mg, dt, ok)
					t.Fail()
				}
				// t.Fatalf("%v %v %v\n", mg, dt, ok)
				// si esta ok ejecuto test de lectura
			})
		}
	}
}
