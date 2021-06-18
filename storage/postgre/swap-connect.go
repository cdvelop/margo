package postgre

// SwapConnect intercambia a la conexion por defecto del motor de la base de datos
func (d *DNS) SwapConnect() (backupDNS DNS) {
	backupDNS = *d
	d.DataBaseName = d.DataBasEngine()
	d.UserDatabase = d.DataBasEngine()
	return
}
