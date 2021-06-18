package postgre

import (
	"fmt"
	"log"

	"github.com/cdvelop/margo/storage/config"
)

//DeleleDataBASE ..
func (d *DNS) DeleleDataBASE() {
	dnsBkp := d.SwapConnect()
	db := config.ConnectDB(d) //seteo el pool de conexiones
	defer db.Close()
	sql := fmt.Sprintf("DROP DATABASE IF EXISTS %v;", dnsBkp.DataBaseName)
	if _, err := db.Exec(sql); err != nil {
		log.Fatalf("error %v al eliminar base de datos sql: %v ", err, sql)
	}
	*d = dnsBkp
	fmt.Printf(">>> Base de datos [%v] Eliminada !!!\n", dnsBkp.DataBaseName)
}

func (d *DNS) DeleteRolDataBase() {
	dnsBkp := d.SwapConnect()
	db := config.ConnectDB(d) //seteo el pool de conexiones
	defer db.Close()
	sql := fmt.Sprintf("DROP ROLE IF EXISTS %v;", dnsBkp.UserDatabase)
	if _, err := db.Exec(sql); err != nil {
		log.Fatalf("error %v al eliminar rol usuario base de datos sql: %v ", err, sql)
	}
	*d = dnsBkp
	fmt.Printf(">>> Usuario db Rol [%v] Eliminado !!!\n", dnsBkp.DataBaseName)
}
