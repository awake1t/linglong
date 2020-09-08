package plugins

import (
	"github.com/stacktitan/smb/smb"
)



func ScanSmb(ip string, port string, username string, password string) (err error, result bool) {
	result = false
	//fmt.Println("run smcd")
	//port,_ = strconv.Atoi(port)

	options := smb.Options{
		Host:        ip,
		Port:        445,
		User:        username,
		Password:    password,
		Domain:      "",
		Workstation: "",
	}

	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			result = true
		}
	}
	return err, result
}

//
//debug := false
//session, err := smb.NewSession(options, debug)
//if err != nil {
//log.Fatalln("[!]", err)
//}
//defer session.Close()
//
//if session.IsSigningRequired {
//log.Println("[-] Signing is required")
//} else {
//log.Println("[+] Signing is NOT required")
//}
//
//if session.IsAuthenticated {
//log.Println("[+] Login successful")
//} else {
//log.Println("[-] Login failed")
//}
//
//if err != nil {
//log.Fatalln("[!]", err)
//}