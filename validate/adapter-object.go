package validate

// Object input to validate
type Object interface {
	TableName() string
	FieldValue() map[string]string //nombre de los campos y sus valores
}
