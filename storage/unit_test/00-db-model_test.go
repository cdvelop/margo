package unit_test

import (
	"github.com/cdvelop/margo/model"
	"github.com/cdvelop/margo/storage/migrate"
	"github.com/cdvelop/margo/ui/artifac/input"
)

var allTables []migrate.Table

func init() {
	text := input.Text{}
	radio_genero := input.NewRadio("Dama", "Varon")

	t1 := model.Table{
		Noun: "usuario",
		Fields: []model.Field{
			{Name: "id_usuario", DataType: "TEXT"},
			{Name: "nombre", DataType: "TEXT", Input: text, Validate: true},
			{Name: "apellido", DataType: "TEXT"},
			{Name: "genero", DataType: "TEXT", Input: radio_genero, Validate: true},
		},
	}
	t1.AddTableToIndex()

	t2 := model.Table{
		Noun: "especialidad",
		Fields: []model.Field{
			{Name: "id_especialidad", DataType: "TEXT"},
			{Name: "nombre_especialidad", DataType: "TEXT"},
			{Name: "id_usuario", DataType: "TEXT"},
			{Name: "detalle", DataType: "TEXT"},
		},
	}
	t1.AddTableToIndex()

	allTables = append(allTables, t1, t2)

}
