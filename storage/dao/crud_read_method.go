package dao

import (
	"fmt"
	"log"

	"github.com/cdvelop/margo/storage/config"
	"github.com/cdvelop/margo/storage/share"
)

//QueryWithoutANSWER SinResultado ejecuta sql en bd con sin respuesta de mas de 1 operacion un error
//recibe sql y mensaje a mostrar en consola
func QueryWithoutANSWER(sql, mje string, da DbAdapter) bool { //20.5.12.19:30
	db := config.Open(da)
	_, err := db.Exec(sql)
	if err != nil {
		log.Printf("%v %v", err, sql)
		return false
	}
	fmt.Println(mje)
	return true
}

//QueryOne .
// https://my.oschina.net/nowayout/blog/139398
func QueryOne(sql string, da DbAdapter) (rowMap map[string]string, ok bool) {
	db := config.Open(da)

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
		return
	}

	rowMap, err = share.FetchOne(rows)
	if err != nil {
		log.Println(err)
		return
	}
	ok = true
	return
}

//QueryAll .
func QueryAll(sql string, da DbAdapter) (rowsMap []map[string]string, ok bool) {
	db := config.Open(da)

	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
		return
	}

	rowsMap, err = share.FetchAll(rows)
	if err != nil {
		log.Println(err)
		return
	}
	ok = true
	return
}
