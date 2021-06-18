package sqlite

import (
	// sqlite3 .
	_ "github.com/mattn/go-sqlite3"
)

//DNS formato cadena conexion
type DNS struct {
	DataBaseName string //nombre de la base de datos
}

func NewConnection(dataBaseName string) *DNS {
	d := DNS{
		DataBaseName: dataBaseName,
	}
	return &d
}
