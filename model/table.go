package model

type Table struct {
	Noun   string  //nombre de la tabla
	Fields []Field //campos de la tabla
}

func (t Table) Name() string {
	return t.Noun
}

func (t Table) FieldsDataType() []map[string]string {
	var allField []map[string]string
	for _, prop := range t.Fields {
		fieldOut := map[string]string{prop.Name: prop.DataType}
		allField = append(allField, fieldOut)
	}
	return allField
}
