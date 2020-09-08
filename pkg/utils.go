package pkg

import (
	"crypto/md5"
	"fmt"
	"io"
	"linglong/models"
	"sync"
)

var (
	aliveIpList []models.IpAddr
	mutex       sync.Mutex
)

func init() {
	aliveIpList = make([]models.IpAddr, 0)
}

func MD5(s string) (m string) {
	h := md5.New()
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func MakeTaskHash(k string) string {
	hash := MD5(k)
	return hash
}
