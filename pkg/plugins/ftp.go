package plugins

import (
	"github.com/jlaffaye/ftp"
	"time"
)

func ScanFtp(ip string, port string, username string, password string) (err error, result bool) {
	conn, err := ftp.DialTimeout(ip+":"+port, time.Second*3)
	if err == nil {
		err = conn.Login(username, password)
		if err == nil {
			result = true
			conn.Logout()
		}
	}
	return err, result
}
