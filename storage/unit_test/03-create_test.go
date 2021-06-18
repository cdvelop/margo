package unit_test

import (
	"fmt"
	"testing"

	"github.com/cdvelop/margo/storage/dao"
)

func Create(db dao.CRUDadapter, t *testing.T) {
	for prueba, data := range dataTestCRUD {
		t.Run((prueba + " " + fmt.Sprint(data.result)), func(t *testing.T) {
			mg, dt, ok := dao.CREATE(data.object, db)
			if ok != data.result {
				t.Fatalf("%v %v %v\n", mg, dt, ok)
			}
			if ok {
				// si esta ok ejecuto test de lectura
				// t.Logf("id: %v resp db [%v]\n", mg, ok)
				objRead := dataTestCRUD[prueba]
				objRead.idRecovered = mg
				dataTestCRUD[prueba] = objRead
			}
		})
	}
}
