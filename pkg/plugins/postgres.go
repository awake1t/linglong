package plugins

import (
	_ "github.com/lib/pq"

	"database/sql"
	"fmt"
)

func ScanPostgres(ip string, port string, username string, password string) (err error, result bool) {
	//fmt.Println( fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", username, password, ip, port, "postgres", "disable"))
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", username, password, ip, port, "postgres", "disable"))
	if err == nil {
		defer db.Close()
		err = db.Ping()
		if err == nil {
			result = true
		}
	}
	return err, result
}
