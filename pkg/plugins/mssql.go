package plugins

import (
	"database/sql"
	_ "github.com/denisenkom/go-mssqldb"
)

func ScanMssql(ip string, port string, username string, password string) (err error, result bool) {
	db, err := sql.Open("mssql", "server="+ip+";port="+port+";user id="+username+";password="+password+";database=master")
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result = true
		}
	}
	return err, result
}
