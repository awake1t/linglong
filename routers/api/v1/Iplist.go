package v1

import (
	"bytes"
	"context"
	"fmt"
	"github.com/Ullaakut/nmap"
	"github.com/astaxie/beego/validation"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/malfunkt/iprange"
	"io/ioutil"
	"linglong/models"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"linglong/routers/tools/masscan"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

var (
	cgf = g.Client()
)

func MergeUrl(url, tmp string) (resurl bytes.Buffer) {
	resurl.WriteString(url)
	resurl.WriteString(tmp)
	return resurl
}

func GetIplist(c *gin.Context) {

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	code := e.SUCCESS

	data["lists"] = models.GetIplist(utils.GetPage(c), 10, maps,"")
	data["total"] = models.GetIplistTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//查询
func GetIplistSearch(c *gin.Context) {
	code := e.INVALID_PARAMS
	valid := validation.Validation{}
	data := make(map[string]interface{})
	maps := make(map[string]interface{})

	ip := c.Query("ip")
	port := c.Query("port")
	title := c.Query("title")
	if ip != "" {
		maps["ip"] = ip
	}
	if port != "" {
		maps["port"] = port
	}


	if ! valid.HasErrors() {
		data["lists"] = models.GetIplist(utils.GetPage(c), 10, maps,title)
		data["total"] = models.GetIplistTotal(maps)
		code = e.SUCCESS
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

func AddIplist(c *gin.Context) {

	start := time.Now()
	massRes := MasscanStart()

	fmt.Println("massRes:", massRes)
	costTime := time.Since(start)
	fmt.Println("costTime:", costTime)

	data := make(map[string]interface{})
	data["task_type"] = "masscan"
	data["succes_num"] = len(massRes)
	data["run_time"] = fmt.Sprintf("%s", costTime)
	data["status"] = 1
	models.AddLog(data)

	wg := &sync.WaitGroup{}
	// 创建一个buffer为thread * 2的channel
	thread := 50
	taskChan := make(chan MassScanRes, 50*2)

	// 创建Thread个协程
	for i := 0; i < thread; i++ {
		go NmapScan(taskChan, wg)
	}

	for _, task := range massRes {
		wg.Add(1)
		taskChan <- task
	}
	// 生产完成后，从生产方关闭task
	close(taskChan)

	wg.Wait()

}


func InitMasscan() {

	start := time.Now()
	massRes := MasscanStart()

	costTime := time.Since(start)
	data := make(map[string]interface{})
	data["taskid"] = 0
	data["task_name"] = "masscan"
	data["task_type"] = "masscan"
	data["all_num"] = len(massRes)
	data["succes_num"] = 0
	data["run_time"] = fmt.Sprintf("%s", costTime)
	data["error"] = "nil"
	data["status"] = 1
	models.AddLog(data)

	fmt.Println("massRes:", massRes)

	// 并发处理masscan扫描结果
	wg := &sync.WaitGroup{}
	// 创建一个buffer为thread * 2的channel
	thread := 2
	taskChan := make(chan MassScanRes, 50*2)

	// 创建Thread个协程
	for i := 0; i < thread; i++ {
		go NmapScan(taskChan, wg)
	}

	for _, task := range massRes {
		wg.Add(1)
		taskChan <- task
	}
	// 生产完成后，从生产方关闭task
	close(taskChan)
	wg.Wait()


	for _,v := range massRes{
		CheckUrlStatus(v)
	}


	fmt.Println("title识别完成,准备登陆路径扫描识别:")

	findLoginWeb()

	fmt.Println("title识别完成，准备删除过期资产")

	masssettings := models.GetSettingTitle()
	expirteCount := masssettings[0].MasscanDeltime
	expirteData := models.GetLogUpdate()
	if len(expirteData) > expirteCount-1{
		delTime := strings.Replace(expirteData[expirteCount-1].CreatedTime, "T", " ", 1)
		models.DelIplistUpdate(delTime[:19])
	}

}

//写入用户输入ip range 类型。 到需要扫描txt
func writeIpToFile(ipList, filename string) error {
	list, err := iprange.ParseList(ipList)
	if err != nil {
		log.Printf("IP Input error: %s", err)
		return err
	}
	rng := list.Expand()

	fd, _ := os.OpenFile(filename, os.O_TRUNC|os.O_WRONLY, 0644)
	for _, ip := range rng {
		var buf []byte
		ips := fmt.Sprintf("%s\n", ip)
		buf = []byte(ips)
		fd.Write(buf)
	}
	fd.Close()
	return nil
}

type MassScanRes struct {
	Ip       string
	Port     string
	Protocol string
}

func MasscanStart() (massRes []MassScanRes) {

	masssetting := models.GetSettingTitle()

	rate := strconv.Itoa(masssetting[0].MasscanThred)

	massRes = make([]MassScanRes, 0)

	MasscanIps := strings.Split(masssetting[0].MasscanIp, "\n")
	MasscanIpNoScans := strings.Split(masssetting[0].MasscanWhite, "\n")

	var ips bytes.Buffer
	for _, ip := range MasscanIps {
		ips.WriteString(ip)
		ips.WriteString(" ")
	}

	var ipsNoScan bytes.Buffer
	if len(MasscanIpNoScans) < 5 {
		ipsNoScan.WriteString("189.198.198.198")
	} else {
		for _, ip := range MasscanIpNoScans {
			ipsNoScan.WriteString(ip)
			ipsNoScan.WriteString(" ")
		}
	}

	m := masscan.New()

	m.SetSystemPath("masscan")
	m.SetRate(rate)
	args := []string{
		//"--range",s.Conf.Masscan.Rate,
		ips.String(),
		"-p", masssetting[0].MasscanPort,
		"--exclude", ipsNoScan.String(),
	}
	m.SetArgs(args...)

	err := m.Run()
	if err != nil {
		fmt.Println(err)
	}

	// 解析扫描结果
	results, err := m.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Masscan本次扫描到%d个ip\n", len(results))

	for _, result := range results {
		fmt.Println("awake res :", result)
		fmt.Println("awake :' , Protocol: result.Ports[0].Protocol", result.Ports[0].Protocol)
		massResTmp := MassScanRes{Ip: result.Address.Addr, Port: result.Ports[0].Portid}
		massRes = append(massRes, massResTmp)
	}

	fmt.Println("mass去重后结果", massRes)

	return
}

// 读取mass扫描结果，跟当前数据库对比。如果不存在，就namp识别端口入库
func NmapScan(taskChan chan MassScanRes, wg *sync.WaitGroup) {
	data := make(map[string]interface{})
	dataUpdate := make(map[string]interface{})
	for target := range taskChan {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("awake :", err)
				wg.Done()
			}
		}()

		data["ip"] = target.Ip
		data["port"] = target.Port
		//扫描结果入库前对比
		ok, id := models.ExistIplist(target.Ip, target.Port)
		if ok {
			fmt.Println(target.Ip, target.Port, "存在数据库中 更新时间")
			nowTime := time.Now().Format("20060102150405")
			dataUpdate["updated_time"] = nowTime
			models.EditIplist(id, dataUpdate)
			wg.Done()
		} else {
			fmt.Println(target.Ip, target.Port, "不存在数据库中,nmap识别端口指纹，入库")
			ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Minute)
			defer cancel()
			scanner, err := nmap.NewScanner(
				nmap.WithTargets(target.Ip),
				//map.WithTargets("google.com", "facebook.com", "youtube.com"),
				nmap.WithPorts(target.Port),
				nmap.WithContext(ctx),
				nmap.WithSkipHostDiscovery(), // s.args = append(s.args, "-Pn") 加上 -Pn 就不去ping主机，因为有的主机防止ping,增加准确度
			)
			if err != nil {
				log.Fatalf("unable to create nmap scanner: %v", err)
			}

			result, warnings, err := scanner.Run()
			if err != nil {
				log.Fatalf("Unable to run nmap scan: %v", err)
			}

			if warnings != nil {
				log.Printf("Warnings: \n %v", warnings)
			}

			var protocol string
			// Use the results to print an example output
			for _, host := range result.Hosts {
				if len(host.Ports) == 0 || len(host.Addresses) == 0 {
					continue
				}

				for _, port := range host.Ports {
					if port.State.State == "open" {
						b := strconv.Itoa(int(port.ID))
						c := string(b)
						fmt.Println(host.Addresses[0].String(), c, port.Service.Name)
						protocol = port.Service.Name
						wg.Done()
					} else {
						wg.Done()
					}
				}
			}
			data["protocol"] = protocol
			models.AddIplist(data)
		}

	}

}

//
func CheckUrlStatus(target MassScanRes) {
		defer func() {
			if err := recover(); err != nil {
				//更新数据库
				fmt.Println("CheckUrlStatus http错误:", target.Ip,target.Port)
				//UpDataCheckStatus(src, "", "", "", 0, 0)
			}
		}()

		if target.Port == "22" || target.Port == "21" || target.Port == "23" || target.Port == "139" || target.Port == "445" || target.Port == "1433" || target.Port == "3306" || target.Port == "6379" || target.Port == "3389" {
			fmt.Println("常用端口退出:", target.Ip)
			return
		}
		var title string

		var src bytes.Buffer
		src.WriteString(target.Ip)
		src.WriteString(":")
		src.WriteString(target.Port)

		cgf.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
		var tmpurl bytes.Buffer
		if target.Port == "443"{
			tmpurl = MergeUrl("https://", src.String())
		}else{
			tmpurl = MergeUrl("http://", src.String())
		}
		resp, err := cgf.Timeout(time.Second * 3).Get(tmpurl.String())

		if err == nil {
			body, _ := ioutil.ReadAll(resp.Body)
			pTitle := regexp.MustCompile(`(?i:)<title>(.*?)</title>`)
			titleArr := pTitle.FindStringSubmatch(string(body))
			fmt.Println("titleArr", titleArr, target.Ip)
			if titleArr != nil {
				if len(titleArr) == 2 {
					sTitle := titleArr[1]
					if !utf8.ValidString(sTitle) {
						sTitle = mahonia.NewDecoder("gb18030").ConvertString(sTitle)
					}
					title = sTitle
				} else {
					title = "Null"
				}
			} else {
				title = "Null"
			}
		}
		fmt.Println("https:titil update", target.Ip, target.Port, resp.Request.URL)
		dataUpdate := make(map[string]interface{})
		dataUpdate["title"] = title
		dataUpdate["loginurl"] = resp.Request.URL.String()
		models.EditIplistByIp(target.Ip, target.Port, dataUpdate)

	}


//根据关键字判断登陆后台
func findLoginWeb() {
	setting := models.GetSettingTitle()
	iplistTitle := models.GetIplistTitle()

	for _, settingKey := range setting {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("setting.setting Protocol11111x:", err)
			}
		}()

		for _, iplist := range iplistTitle {
			titlekey := strings.Split(settingKey.LoginWord, "\n")
			for _, v := range titlekey {
				if v == "" {
					continue
				}
				if strings.Contains(iplist.Title, v) {
					ok, id := models.ExistWebloginlist(iplist.Ip, iplist.Port)
					if ok {
						nowTime := time.Now().Format("20060102150405")
						dataUpdate := make(map[string]interface{})
						dataUpdate["updated_time"] = nowTime
						models.EditWebloginlist(id, dataUpdate)
					} else {
						defer func() {
							if err := recover(); err != nil {
								fmt.Println("错误:识别敏感关键后台", err)
							}
						}()

						dataloginlist := make(map[string]interface{})
						dataloginlist["ip"] = iplist.Ip
						dataloginlist["port"] = iplist.Port
						dataloginlist["protocol"] = iplist.Protocol
						dataloginlist["title"] = iplist.Title
						dataloginlist["url"] = iplist.Loginurl
						models.AddWebloginlist(dataloginlist)
					}

				}
			}
		}

		//取出资产列表里所有的loginurl，去掉空的
		for _, iplist := range iplistTitle {
			titlekey := strings.Split(settingKey.LoginUrl, "\n")
			for _, v := range titlekey {
				if v == "" {
					continue
				}
				if strings.Contains(iplist.Loginurl, v) {
					ok, id := models.ExistWebloginlist(iplist.Ip, iplist.Port)
					if ok {
						nowTime := time.Now().Format("20060102150405")
						dataUpdate := make(map[string]interface{})
						dataUpdate["updated_time"] = nowTime
						models.EditWebloginlist(id, dataUpdate)
					} else {
						defer func() {
							if err := recover(); err != nil {
								fmt.Println("错误:识别敏感关键url", err)
							}
						}()

						dataloginlist := make(map[string]interface{})
						dataloginlist["ip"] = iplist.Ip
						dataloginlist["port"] = iplist.Port
						dataloginlist["protocol"] = iplist.Protocol
						dataloginlist["title"] = iplist.Title
						dataloginlist["url"] = iplist.Loginurl
						models.AddWebloginlist(dataloginlist)
					}

				}
			}
		}

	}

}
