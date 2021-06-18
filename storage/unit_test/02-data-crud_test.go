package unit_test

import (
	"github.com/cdvelop/margo/model"
)

type kv map[string]string

type dataModel struct {
	object      model.Object
	result      bool
	idRecovered string
}

var (
	defaulTableName = "usuario"
	dataTestCRUD    = map[string]dataModel{

		"Luis campos correctos?": {model.Object{
			Table_Name:  defaulTableName,
			Field_Value: kv{"nombre": "Luis", "apellido": "de las carmenes", "genero": "V"},
		}, true, ""},

		"Maria campos correctos?": {model.Object{
			Table_Name:  defaulTableName,
			Field_Value: kv{"nombre": "Maria", "apellido": "Ruiz", "genero": "D"},
		}, true, ""},

		"genero H existe?": {model.Object{
			Table_Name:  defaulTableName,
			Field_Value: kv{"nombre": "Luis", "apellido": "de las carmenes", "genero": "H"},
		}, false, ""},

		"apellido numerico, se requiere validacion?": {model.Object{
			Table_Name:  defaulTableName,
			Field_Value: kv{"nombre": "Luis", "apellido": "2", "genero": "H"},
		}, false, ""},

		"nombre corresponde a solo texto?": {model.Object{
			Table_Name:  defaulTableName,
			Field_Value: kv{"nombre": "mar1a", "apellido": "de las carmenes"},
		}, false, ""},

		"todos los campos?": {model.Object{
			Table_Name:  defaulTableName,
			Field_Value: kv{"nombre": "Maria"},
		}, false, ""},
	}
)
