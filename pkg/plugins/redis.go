package plugins

import (
	"github.com/go-redis/redis"
	"time"
)

func ScanRedis(ip string, port string, username string, password string) (err error, result bool) {
	client := redis.NewClient(&redis.Options{Addr: ip + ":" + port, Password: password, DB: 0, DialTimeout: time.Second * 3})
	defer client.Close()
	_, err = client.Ping().Result()
	if err == nil {
		result = true
	}
	return err, result
}
