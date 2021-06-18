package model

type Object struct {
	Table_Name  string
	Field_Value map[string]string
}

func (o Object) TableName() string {
	return o.Table_Name
}

func (o Object) FieldValue() map[string]string {
	return o.Field_Value
}

func (o Object) GetModelObject() *map[string]Field {
	return GetTableModelByName(o.Table_Name)
}
