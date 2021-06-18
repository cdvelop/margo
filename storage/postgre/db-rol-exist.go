package postgre

import (
	"fmt"

	"github.com/cdvelop/margo/storage/share"
)

//ExistDataBase verifica si existe una base de datos posgres determinada
func (d *DNS) ExistDataBase() bool {
	sql := fmt.Sprintf(d.DBExists(), d.DataBaseName)
	// log.Printf("sql %v", sql)
	return share.Exists("base de datos", d.DataBaseName, sql, d)
}

//ExistDataBaseROL verifica  si rol usuario aplicacion existe
func (d *DNS) ExistDataBaseROL(rol string) bool {
	sql := fmt.Sprintf(d.ROLExists(), rol)

	return share.Exists("rol", d.UserDatabase, sql, d)
}
