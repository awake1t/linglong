package plugins

import (
	"gopkg.in/mgo.v2"
	"time"
)

//mongodb未授权和弱口令漏洞
func ScanMongodb(ip string, port string, username string, password string) (err error, result bool) {
	session, err := mgo.DialWithTimeout("mongodb://"+username+":"+password+"@"+ip+":"+port+"/"+"admin", time.Second*3)
	if err == nil && session.Ping() == nil {
		defer session.Close()
		if err == nil && session.Run("serverStatus", nil) == nil {
			result = true
		}
	}
	return err, result
}

func MongoUnauth(ip string, port string) (err error, result bool) {
	session, err := mgo.Dial(ip + ":" + port)
	if err == nil && session.Run("serverStatus", nil) == nil {
		result = true
	}
	return err, result
}
