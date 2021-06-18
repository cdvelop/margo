package postgre

import (
	// postgres data base
	_ "github.com/lib/pq"
)

//DNS formato cadena conexion
type DNS struct {
	DataBaseName     string //nombre de la base de datos
	IPLocalServer    string //ip servidor donde estara la base de datos
	DataBasePORT     string //puerto
	UserDatabase     string //usuario base de datos
	PasswordDatabase string //contraseña
}

func NewConnection(userDatabase, passwordDatabase, iPLocalServer, dataBasePORT, dataBaseName string) *DNS {
	d := DNS{
		DataBaseName:     dataBaseName,
		IPLocalServer:    iPLocalServer,
		DataBasePORT:     dataBasePORT,
		UserDatabase:     userDatabase,
		PasswordDatabase: passwordDatabase,
	}
	return &d
}
