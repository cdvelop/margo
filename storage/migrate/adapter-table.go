package migrate

type Table interface {
	Name() string
	FieldsDataType() []map[string]string
}
