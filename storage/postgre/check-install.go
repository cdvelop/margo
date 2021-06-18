package postgre

import (
	"fmt"
	"log"

	"github.com/cdvelop/margo/storage/config"
	"github.com/cdvelop/margo/storage/dao"
)

//CHECK verifica el estado de la base motor db POSTGRES
func (d *DNS) CHECK() bool {
	fmt.Printf(">>> *** Inicio Verificacion de Base de datos [%v] ..... *** <<<\n\n", d.DataBaseName)
	if ready := d.ExistDataBase(); ready { //existe db app?
		return true
	} else {
		//obtengo la conexion actual
		bkpDns := d.SwapConnect()
		defer config.ConnectDB(d).Close()

		if ready = d.ExistDataBase(); ready { //existe postgres?

			if ready = d.ExistDataBaseROL(bkpDns.UserDatabase); !ready { //verificar rol usuario app
				//crear rol app
				if ready = dao.QueryWithoutANSWER(`CREATE USER `+bkpDns.UserDatabase+` PASSWORD '`+d.PasswordDatabase+`';`, ">>> creando rol db", d); !ready {
					log.Fatalf("!!! error al crear rol: %v", bkpDns.UserDatabase)
				}
				fmt.Printf(">>> rol: %v creado\n", bkpDns.UserDatabase)
			}

			if dao.QueryWithoutANSWER(`CREATE DATABASE `+bkpDns.DataBaseName+` OWNER=`+bkpDns.UserDatabase+`;`, ">>> creando db", d) {
				fmt.Printf(">>> base de datos: %v creada\n", bkpDns.DataBaseName)
				// cambiar clave administrador postgres****
				if dao.QueryWithoutANSWER(`ALTER USER postgres WITH PASSWORD '`+d.PasswordDatabase+`';`, ">>> cambiando password db", d) {
					fmt.Print(">>> password postgres actualizada\n")
				}
				if dao.QueryWithoutANSWER(`GRANT ALL PRIVILEGES ON DATABASE `+bkpDns.DataBaseName+` TO `+bkpDns.UserDatabase+`;`, ">>> otorgando privilegios a db", d) {
					fmt.Print(">>> privilegios rol " + bkpDns.UserDatabase + " otorgados a db\n")

					//cambiar a nuevo dns inicial
					*d = bkpDns
					config.ConnectDB(d) //seteo el pool de conexiones
					fmt.Printf(">>> *** Verificacion de Base de datos [%v] Finalizada *** <<<\n\n", d.DataBaseName)
					return true
				}
			}

		} else {
			log.Fatalf("!!! error motor de base de datos %v no existe chequear instalación", d.DataBasEngine())
		}
	}
	return false
}
