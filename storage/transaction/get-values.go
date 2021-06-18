package transaction

import (
	"context"
	"database/sql"
)

//getValuesWithContext retorna valor de una consulta sql
func GetValuesWithContext(sql string, db *sql.DB, ctx context.Context) (out map[string]string, ok bool) {
	out = make(map[string]string) //inicializamos la salida

	row := db.QueryRowContext(ctx, sql)
	var value string
	err := row.Scan(&value)
	if err != nil {
		out["err"] = err.Error()
		return
	}
	out["value"] = value
	ok = true
	return
}
