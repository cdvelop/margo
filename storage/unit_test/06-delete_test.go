package unit_test

import (
	"testing"

	"github.com/cdvelop/margo/model"
	"github.com/cdvelop/margo/storage/dao"
)

func Delete(db dao.CRUDadapter, t *testing.T) {

	for _, dataRead := range dataTestCRUD {
		if dataRead.result { //solo los casos de exito
			deleteObject := model.Object{
				Table_Name: defaulTableName,
				Field_Value: map[string]string{
					"id_usuario": dataRead.idRecovered},
			}

			t.Run(("DELETE: " + dataRead.idRecovered), func(t *testing.T) {
				mg, dt, ok := dao.DELETE(deleteObject, db)
				if !ok {
					t.Logf("message: %v data %v ok[%v]\n", mg, dt, ok)
					t.Fail()
				}
			})
		}
	}
}
