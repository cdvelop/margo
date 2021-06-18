package model

// central tables index
// índice central de tablas para buscar sus propiedades por texto

var (
	//nombre tabla / campo /propiedades
	globaLindex map[string]map[string]Field
)

func init() {
	globaLindex = make(map[string]map[string]Field)
}

//agrega la tabla al indice general del sistema
func (table Table) AddTableToIndex() {
	newFields := make(map[string]Field)

	for _, fieldIn := range table.Fields {
		newFields[fieldIn.Name] = fieldIn
	}

	globaLindex[table.Noun] = newFields
}

func GetTableModelByName(tableName string) *map[string]Field {
	table := make(map[string]Field)
	if tableFound, ok := globaLindex[tableName]; ok {
		table = tableFound
	}
	return &table
}
