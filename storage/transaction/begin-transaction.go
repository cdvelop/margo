package transaction

import (
	"context"
	"database/sql"
	"log"

	"github.com/cdvelop/margo/storage/config"
)

// BeginTX inicia una nueva transacción. y contexto
func BeginTX(db *config.Connection) (*sql.Tx, context.Context) {
	cx := context.Background()
	tx, err := db.BeginTx(cx, nil)

	if err != nil {
		log.Println(err)
		return nil, nil
	}
	return tx, cx
}
