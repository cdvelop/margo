package unit_test

import (
	"os"
	"testing"

	"github.com/cdvelop/margo/storage/migrate"
	"github.com/cdvelop/margo/storage/postgre"
	"github.com/cdvelop/margo/storage/sqlite"
)

var db_pg = postgre.NewConnection("test", "1", "127.0.0.1", "5432", "test")

func Test_PostgresCreateDB(t *testing.T) {
	//eliminar base de datos
	db_pg.DeleleDataBASE()
	//borrar usuario
	db_pg.DeleteRolDataBase()

	if !db_pg.CHECK() {
		t.Fail()
	}

	if !migrate.CreateAllTABLES(db_pg, allTables...) {
		t.Fail()
	}

}

var db_sqlite = sqlite.NewConnection("test")

func Test_SqliteCreateDB(t *testing.T) {
	//eliminar base de datos
	os.Remove("./test.db")

	if !migrate.CreateAllTABLES(db_sqlite, allTables...) {
		t.Fail()
	}

}
