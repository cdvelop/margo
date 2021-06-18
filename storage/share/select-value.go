package share

import "github.com/cdvelop/margo/storage/config"

//SelectValue retorna valor de una consulta sql
func SelectValue(sql string, d DbAdapter) (out string, ok bool) {
	db := config.Open(d)
	row := db.QueryRow(sql)
	err := row.Scan(&out)
	if err != nil {
		out = err.Error()
		return
	}
	ok = true
	return
}
