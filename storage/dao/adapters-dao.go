package dao

type DbAdapter interface {
	DataBasEngine() string    //motor base de datos postgres,mysql,sqlite3 etc
	ConnectionString() string //cadena con formato de conexion base de datos dns
}

//ORM soporte para distintas db
type ORM interface {
	SQLsyntax() string //ej postgres:"$1", sqlite: "?"
	// "%v=?", k
	SetListSyntax(string, byte, *[]string)
	TotalValuesSyntax(map[string]string) string //"?" or "$1" idjugador = ? etc
	// SQLParameters string  //ej: postgres: "$" sqlite "?",
	MakeSqInsertSyntax(*byte, *[]string)
}

type CRUDadapter interface {
	DbAdapter
	ORM
}

// RequestObject solicitud
type RequestObject interface {
	TableName() string
	FieldValue() map[string]string //nombre del campo y su valor
}
