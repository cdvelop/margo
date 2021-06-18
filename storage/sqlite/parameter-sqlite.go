package sqlite

import "fmt"

func (dns *DNS) DataBasEngine() string {
	return "sqlite3"
}

// ConnectionString formato cadena de conexion
func (dns *DNS) ConnectionString() string {
	return fmt.Sprintf("%v.db", dns.DataBaseName)
}

func (d *DNS) SQLTableInfo() string {
	return "SELECT name FROM PRAGMA_TABLE_INFO('%v');"
}

func (d *DNS) SQLColumName() string {
	return "name"
}

func (d *DNS) SQLDropTable() string {
	return "DROP TABLE IF EXISTS %v;" //sql de eliminacion de tabla
}
