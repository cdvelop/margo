package share

type DbAdapter interface {
	DataBasEngine() string    //motor base de datos postgres,mysql,sqlite3 etc
	ConnectionString() string //cadena con formato de conexion base de datos dns
}
