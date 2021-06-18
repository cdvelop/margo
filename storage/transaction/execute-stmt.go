package transaction

import (
	"context"
	"database/sql"
	"log"
)

//ExecuteSQLStatement ingresa un objeto en db
func ExecuteSQLStatement(sql string, arg []interface{}, tx *sql.Tx, ctx context.Context) (ok bool) {
	stmt, err := tx.PrepareContext(ctx, sql)
	if err != nil {
		log.Println(err)
		return
	}
	defer stmt.Close()

	// Exec.
	r, err := stmt.ExecContext(ctx, arg...)
	if err != nil {
		log.Println(err)
		return
	}

	i, _ := r.RowsAffected()
	if i != 1 {
		log.Println("err se esperaba solo una fila afectada")
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Println(err)
		return
	}

	ok = true
	return
}
