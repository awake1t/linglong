package brute

import (
	"fmt"
	"github.com/fatih/color"
	"linglong/models"
	"linglong/pkg"
	"linglong/pkg/plugins"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	mutex       sync.Mutex
	successHash map[string]bool
	bruteResult map[string]models.Service
)

func saveRes(target models.Service, h ,taskruntime string ,taskid int  ) {

	setTaskHask(h)
	_, ok := bruteResult[h]
	if !ok {
		mutex.Lock()
		//爆破结果写入数据库中
		color.Cyan("[+] %s %d %s %s \n", target.Ip, target.Port, target.UserName, target.PassWord)
		data := make(map[string]interface{})
		data["ip"] = target.Ip
		data["port"] = target.Port
		data["protocol"] = target.Protocol
		data["user"] = target.UserName
		data["pass"] = target.PassWord
		data["taskid"] = taskid
		//data["vulntype"] = protocol // todo:漏洞类型要判断下
		//data["task_id"] = protocol
		data["task_time"] = taskruntime
		models.AddPortBruteRes(data)

		bruteResult[h] = models.Service{Ip: target.Ip, Port: target.Port, Protocol: target.Protocol, UserName: target.UserName, PassWord: target.PassWord}
		mutex.Unlock()
	}
}

// 消费者 每个协程都从生产者的channel中读取数据后，开启扫描
func runBrute(taskChan chan models.Service,taskid int, taskruntime string, wg *sync.WaitGroup) {

	for target := range taskChan {
		//fmt.Println("now is :runBrute ",target)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("错误:runBrute error", target,err)
			}
		}()

		protocol := strings.ToUpper(target.Protocol)

		var k string
		if protocol == "REDIS" || protocol == "FTP" || protocol == "SNMP" || protocol == "POSTGRESQL" || protocol == "SSH" {
			k = fmt.Sprintf("%v-%v-%v", target.Ip, target.Port, target.Protocol)
		} else {
			k = fmt.Sprintf("%v-%v-%v", target.Ip, target.Port, target.UserName)
		}

		// 生成唯一hask
		h := pkg.MakeTaskHash(k)

		if checkTashHash(h) {
			wg.Done()
			continue
		}
		fmt.Fprintf(os.Stdout, "Now is %s %s %s\r", target.Ip,target.UserName, target.PassWord)
		err, res := plugins.ScanFuncMap[protocol](target.Ip, strconv.Itoa(target.Port), target.UserName, target.PassWord)
		if err == nil && res == true {
			saveRes(target,  h,taskruntime,taskid)
		} else {
			//fmt.Println("插件爆破时错误:", err)
		}
		wg.Done()
	}

}

func RunTask(scanTasks []models.Service, thread,taskid int,taskruntime string) {
	start := time.Now()
	wg := &sync.WaitGroup{}
	successHash = make(map[string]bool)
	bruteResult = make(map[string]models.Service)

	// 创建一个buffer为thread * 2的channel
	taskChan := make(chan models.Service, thread*2)

	// 创建Thread个协程
	for i := 0; i < thread; i++ {
		go runBrute(taskChan,taskid,taskruntime, wg)

	}

	// 生产者，不断的把生产要扫描的数据，存放到 channel，直到channel阻塞
	for _, task := range scanTasks {
		wg.Add(1)
		taskChan <- task
	}

	// 生产完成后，从生产方关闭task
	close(taskChan)

	wg.Wait()
	waitTimeout(wg, 3*time.Second)

	color.Red("Scan complete. %d vulnerabilities found! \n", len(bruteResult))
	//end := time.Now()
	costTime := time.Since(start)
	data := make(map[string]interface{})

	// 获取这个任务id扫描成功的数量
	maps := make(map[string]interface{})
	maps["task_id"] = taskid
	maps["task_time"] = taskruntime
	succesnum := models.GetPortBruteResTotal(maps)
	fmt.Println("succesnum:",succesnum)
	fmt.Println(" end.Sub(start):", costTime)

	data["succes_num"] = succesnum
	data["run_time"] = fmt.Sprintf("%s",costTime)
	data["status"] = 2
	models.EditTaskLogTaskId(taskid,data)

}

func WriteToFile(wireteString, filename string) {

	fd, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(wireteString)
	fd.Write(buf)

}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // 超时未响应
	}
}

//func GenerateTaskTmp(addr []models.IpAddr, userList []string) (scanTasks []models.Service) {
//	for _, u := range userList {
//		uk := strings.Split(u,":")
//			for _, ip := range addr {
//				scanTask := models.Service{Ip: ip.Ip, Port: ip.Port, Protocol: ip.Protocol, UserName: uk[0], PassWord: uk[1]}
//				scanTasks = append(scanTasks, scanTask)
//			}
//	}
//	return
//}

func GenerateTask(addr []models.IpAddr, userList []string, passList []string) (scanTasks []models.Service) {
	//每个都生成一个空的账号密码，用于爆破空账号密码
	scanTasks = make([]models.Service, 0)
	for _, ip := range addr {
		if ip.Protocol == "REDIS" || ip.Protocol == "FTP" || ip.Protocol == "POSTGRESQL" || ip.Protocol == "SSH" {
			scanTask := models.Service{Ip: ip.Ip, Port: ip.Port, Protocol: ip.Protocol, UserName: "", PassWord: ""}
			scanTasks = append(scanTasks, scanTask)
		}
	}

	for _, u := range userList {
		for _, p := range passList {
			for _, ip := range addr {
				scanTask := models.Service{Ip: ip.Ip, Port: ip.Port, Protocol: ip.Protocol, UserName: u, PassWord: p}
				scanTasks = append(scanTasks, scanTask)
			}
		}
	}

	return
}

// 标记特定服务的特定用户是否破解成功，成功的话不再尝试破解该用户
//SuccessHash map[string]bool hash唯一
func checkTashHash(hash string) bool {
	_, ok := successHash[hash]
	return ok
}

func setTaskHask(hash string) () {
	mutex.Lock()
	successHash[hash] = true
	mutex.Unlock()
}
