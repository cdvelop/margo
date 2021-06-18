package unit_test

import (
	"testing"

	"github.com/cdvelop/margo/storage/migrate"
)

func Test_CRUDobjectPOSTGRE(t *testing.T) {
	migrate.DeleteDataFromTABLE(defaulTableName, db_pg)

	Create(db_pg, t)

	Read(db_pg, t)

	Update(db_pg, t)

	Delete(db_pg, t)

}

func Test_CRUDobjectSQLITE(t *testing.T) {
	migrate.DeleteDataFromTABLE(defaulTableName, db_sqlite)

	Create(db_sqlite, t)

	Read(db_sqlite, t)

	Update(db_sqlite, t)

	Delete(db_sqlite, t)

}
