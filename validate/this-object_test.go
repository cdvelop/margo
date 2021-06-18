package validate_test

import (
	"testing"

	"github.com/cdvelop/margo/model"
	"github.com/cdvelop/margo/ui/artifac/input"
	"github.com/cdvelop/margo/validate"
)

const (
	defaulTableName = "usuario"
	idKey           = "id_usuario"
	nameKey         = "name"
	genderKey       = "gender"
	descriptionKey  = "description"
)

func init() {

	radio_genero := input.NewRadio("Dama", "Varon")
	text := input.Text{}
	number := input.Number{}

	t1 := model.Table{
		Noun: defaulTableName,
		Fields: []model.Field{
			{Name: idKey, Input: number, Validate: true},
			{Name: nameKey, Input: text, Validate: true},
			{Name: genderKey, Input: radio_genero, Validate: true},
			{Name: descriptionKey, Input: text, Validate: false},
		},
	}

	t1.AddTableToIndex()

}

type kv map[string]string

var (
	datatest = map[string]struct {
		object   model.Object //objeto a testear
		Validate string       //validar todo
		result   bool         //resultado esperado
	}{
		"objeto nuevo ok": {model.Object{Table_Name: defaulTableName,
			Field_Value: kv{nameKey: "Luis", genderKey: "D", descriptionKey: "no tiene"},
		}, "new", true},

		"todos los campos": {model.Object{Table_Name: defaulTableName,
			Field_Value: kv{idKey: "123456", nameKey: "Luis", genderKey: "D", descriptionKey: "no tiene"},
		}, "all", true},

		"new faltan campos": {model.Object{Table_Name: defaulTableName,
			Field_Value: kv{nameKey: "Luis"},
		}, "new", false},

		"no validado numero en vez de nombre": {model.Object{Table_Name: defaulTableName,
			Field_Value: kv{nameKey: "1u50"},
		}, "", false},
	}
)

func Test_Validate(t *testing.T) {
	for prueba, data := range datatest {
		t.Run((prueba), func(t *testing.T) {
			message, ok := validate.This(data.object, data.Validate)
			if ok != data.result {
				t.Fatalf("%v %v", message, ok)
			}
		})
	}
}
