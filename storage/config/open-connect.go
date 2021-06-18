package config

import (
	"database/sql"
	"fmt"
	"sync"
)

//Connection acceso a base de datos
type Connection struct {
	*sql.DB
}

var (
	once sync.Once //se ejecutara solo una vez
	err  error
	pool *sql.DB //config base de datos
)

//Open Abrir devuelve una referencia de base de datos.
func Open(a DbAdapter) *Connection {
	once.Do(func() {
		pool = ConnectDB(a)
	})
	return &Connection{pool}
}

//ConnectDB Conn con base de datos
func ConnectDB(a DbAdapter) *sql.DB {
	// pool, err = sql.Open(dns.DataBasEngine, fmt.Sprintf("%v://%v:%v@%v:%v/%v?sslmode=disable", dns.DataBasEngine, dns.UserDatabase, dns.PasswordDatabase, dns.IPLocalServer, dns.DataBasePORT, dns.DataBaseName))
	pool, err = sql.Open(a.DataBasEngine(), a.ConnectionString())
	if err != nil {
		fmt.Printf("!!! error de conexion %v", err)
		// return
	}
	err = pool.Ping()
	if err != nil {
		fmt.Printf("!!! error ping: %v", err)
		// return
	}
	// alexedwards.net/blog/configuring-sqldb - odb.SetMaxOpenConns(25) - odb.SetMaxIdleConns(25) - odb.SetConnMaxLifetime(5 * time.Minute)
	// log.Printf("conexion dns: %v", dns)
	return pool
}
