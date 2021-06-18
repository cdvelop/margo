package postgre

func (d *DNS) SQLTableInfo() string {
	return "SELECT column_name FROM information_schema.columns WHERE table_name = '%v';"
}

func (d *DNS) SQLColumName() string {
	return "column_name"
}

func (d *DNS) SQLDropTable() string {
	return "DROP TABLE IF EXISTS %v CASCADE;" //sql de eliminacion de tabla
}

func (d *DNS) DBExists() string {
	return "SELECT 1 FROM pg_database WHERE datname='%v';"
}

func (d *DNS) ROLExists() string {
	return "SELECT 1 FROM pg_user WHERE usename = '%v';"
}
