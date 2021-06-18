package unit_test

import (
	"reflect"
	"testing"

	"github.com/cdvelop/margo/model"
	"github.com/cdvelop/margo/storage/dao"
)

func Read(db dao.DbAdapter, t *testing.T) {

	for _, dataRead := range dataTestCRUD {

		if dataRead.result { //solo los casos de exito
			newObject := model.Object{
				Table_Name:  defaulTableName,
				Field_Value: map[string]string{"id_usuario": dataRead.idRecovered},
			}
			createdObject := dataRead.object.Field_Value
			createdObject["id_usuario"] = dataRead.idRecovered

			t.Run(("READ: "), func(t *testing.T) {
				out, ok := dao.READallColumnsByID(db, newObject)
				if !reflect.DeepEqual(out, createdObject) {
					t.Logf("!!! READ data: [%v] resp [%v]\n", out, ok)
					t.Fail()
				}
			})
		}
	}
}
