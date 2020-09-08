package common

import (
	"bufio"
	"fmt"
	"linglong/models"
	"os"
	"strconv"
	"strings"
)
// 读取用户/密码字典
func ReadUserDict(userDict string) (users []string, err error) {
	file, err := os.Open(userDict)
	if err != nil {
		fmt.Println("Open user dict file err, %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		user := strings.TrimSpace(scanner.Text())
		if user != "" {
			users = append(users,user)
		}
	}
	return users, err
}



// 读取ip,port列表
func ReadIpList(iplist string) (ipList []models.IpAddr) {
	//ipListFile, err := os.Open(fileName)
	//if err != nil {
	//	fmt.Printf("Open ip List file err, %v", err)
	//}
	//
	//defer ipListFile.Close()
	//
	//scanner := bufio.NewScanner(ipListFile)
	scanner := bufio.NewScanner(strings.NewReader(iplist))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fmt.Println("line:",line)
		ipPort := strings.TrimSpace(line)
		t := strings.Split(ipPort, ":")
		ip := t[0]
		portProtocol := t[1]
		tmpPort := strings.Split(portProtocol, "|")
		// ip列表中指定了端口对应的服务
		if len(tmpPort) == 2 {
			port, _ := strconv.Atoi(tmpPort[0])
			protocol := strings.ToUpper(tmpPort[1])
			if SupportProtocols[protocol] {
				addr := models.IpAddr{Ip: ip, Port: port, Protocol: protocol}
				ipList = append(ipList, addr)
			} else {
				fmt.Printf("Not support %v, ignore: %v:%v", protocol, ip, port)
			}
		} else {
			// 通过端口查服务
			port, err := strconv.Atoi(tmpPort[0])
			if err == nil {
				protocol, ok := PortNames[port]
				if ok && SupportProtocols[protocol] {
					addr := models.IpAddr{Ip: ip, Port: port, Protocol: protocol}
					ipList = append(ipList, addr)
				}
			}
		}

	}

	return ipList
}
