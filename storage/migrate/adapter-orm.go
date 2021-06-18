package migrate

type ORM interface {
	SQLTableInfo() string //sql como obtiene la base de datos el nombre de la tabla
	SQLDropTable() string //sql de eliminacion de tabla
	SQLColumName() string //sql como se llama a la columna en el motor de base de datos
	// DBExists() string     //consulta si db existe
}

type DbAdapter interface {
	DataBasEngine() string    //motor base de datos postgres,mysql,sqlite3 etc
	ConnectionString() string //cadena con formato de conexion base de datos dns
}
