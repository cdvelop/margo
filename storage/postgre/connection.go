package postgre

import "fmt"

//DataBasEngine "postgres"
func (dns *DNS) DataBasEngine() string {
	return "postgres"
}

// ConnectionString formato cadena de conexion
// postgres:// user : PasswordDatabase  @  127.0.0.1  :  5432 / nombrebasedatos   ?sslmode=disable"
func (dns *DNS) ConnectionString() string {
	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", dns.UserDatabase, dns.PasswordDatabase, dns.IPLocalServer, dns.DataBasePORT, dns.DataBaseName)
}
