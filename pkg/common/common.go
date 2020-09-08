package common

import "time"

func GetNowTome()string{
	return time.Now().Format("20060102150405")
}
