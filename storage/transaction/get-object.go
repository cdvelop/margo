package transaction

import (
	"context"
	"database/sql"

	"github.com/cdvelop/margo/storage/share"
)

//GetObjCTX retorna objeto segun sql
func GetObjCTX(sql *string, db *sql.DB, ctx context.Context) (out map[string]string, ok bool) {
	out = make(map[string]string) //inicializamos la salida

	rows, err := db.QueryContext(ctx, *sql)
	if err != nil {
		out["err"] = err.Error()
		return
	}

	out, err = share.FetchOne(rows)
	if err != nil {
		out["err"] = err.Error()
		return
	}

	ok = true
	return
}
