package v1

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ullaakut/nmap"
	"github.com/axgle/mahonia"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"github.com/malfunkt/iprange"
	"io"
	"io/ioutil"
	"linglong/models"
	"linglong/pkg/common"
	"linglong/pkg/e"
	"linglong/pkg/utils"
	"linglong/routers/api/v1/alyaze"
	"linglong/routers/tools/masscan"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unicode/utf8"
)

var (
	cgf    = g.Client()
	wa     *alyaze.WebAnalyzer
	thread = 5
	err    error

	crawlCount      = 0
	searchSubdomain = false
	redirect        = false

	lock sync.Mutex
)

type XrayResAll struct {
	CreateTime int64 `json:"create_time"`
	Detail     struct {
		Addr     string     `json:"addr"`
		Payload  string     `json:"payload"`
		Snapshot [][]string `json:"snapshot"`
		Extra    struct {
			Author string `json:"author"`
			Param  struct {
			} `json:"param"`
		} `json:"extra"`
	} `json:"detail"`
	Plugin string `json:"plugin"`
	Target struct {
		URL string `json:"url"`
	} `json:"target"`
}

func MergeUrl(url, tmp string) (resurl bytes.Buffer) {
	resurl.WriteString(url)
	resurl.WriteString(tmp)
	return resurl
}

func GetIplist(c *gin.Context) {
	protocol := c.Query("protocol")
	ip := c.Query("ip")
	port := c.Query("port")
	title := c.Query("title")
	finger := c.Query("finger")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if protocol != "" {
		maps["protocol"] = protocol
	}

	if ip != "" {
		maps["ip"] = ip
	}

	if port != "" {
		maps["port"] = port
	}

	if title != "" {
		maps["title"] = title
	}

	if finger != "" {
		maps["finger"] = finger
	}
	code := e.SUCCESS

	data["lists"] = models.GetIplist(utils.GetPage(c), 10, maps)
	data["total"] = models.GetIplistTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//查询
//func GetIplistSearch(c *gin.Context) {
//	code := e.INVALID_PARAMS
//	valid := validation.Validation{}
//	data := make(map[string]interface{})
//	maps := make(map[string]interface{})
//
//	ip := c.Query("ip")
//	port := c.Query("port")
//	title := c.Query("title")
//	if ip != "" {
//		maps["ip"] = ip
//	}
//	if port != "" {
//		maps["port"] = port
//	}
//
//	if ! valid.HasErrors() {
//		data["lists"] = models.GetIplist(utils.GetPage(c), 10, maps, title)
//		data["total"] = models.GetIplistTotal(maps)
//		code = e.SUCCESS
//	} else {
//		for _, err := range valid.Errors {
//			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
//		}
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"code": code,
//		"msg":  e.GetMsg(code),
//		"data": data,
//	})
//}

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
	var HttpRes = []string{}
	//os.Exit(1)

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

	//fmt.Println("massRes:", massRes)

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

	for _, v := range massRes {
		HttpRes = append(HttpRes, CheckUrlStatus(v))
	}
	// 登录页面识别
	findLoginWeb()

	ScanPocXray("", HttpRes, 30)

	CheckFinger(HttpRes)

	masssettings := models.GetSettingTitle()
	expirteCount := masssettings[0].MasscanDeltime
	expirteData := models.GetLogUpdate()
	if len(expirteData) > expirteCount-1 {
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
		fmt.Println("mass run  err : ",err)
	}

	results, err := m.Parse()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Masscan find %d ips \n", len(results))

	for _, result := range results {
		//fmt.Println("awake :' , Protocol: result.Ports[0].Protocol", result.Ports[0].Protocol)
		massResTmp := MassScanRes{Ip: result.Address.Addr, Port: result.Ports[0].Portid}
		massRes = append(massRes, massResTmp)
	}

	fmt.Println("mass dup is ", massRes)

	return
}

// 读取mass扫描结果，跟当前数据库对比。如果不存在，就namp识别端口入库
func NmapScan(taskChan chan MassScanRes, wg *sync.WaitGroup) {
	data := make(map[string]interface{})
	dataUpdate := make(map[string]interface{})
	for target := range taskChan {
		defer func() {
			if err := recover(); err != nil {
				//fmt.Println("awake :", err)
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
func CheckUrlStatus(target MassScanRes) string {
	defer func() {
		if err := recover(); err != nil {
			//更新数据库
			fmt.Println("CheckUrlStatus http错误:", target.Ip, target.Port)
			//UpDataCheckStatus(src, "", "", "", 0, 0)
		}
	}()

	if target.Port == "22" || target.Port == "21" || target.Port == "23" || target.Port == "139" || target.Port == "445" || target.Port == "1433" || target.Port == "3306" || target.Port == "6379" || target.Port == "3389" {
		fmt.Println("常用端口退出:", target.Ip)
		return ""
	}
	var title string

	var src bytes.Buffer
	src.WriteString(target.Ip)
	src.WriteString(":")
	src.WriteString(target.Port)

	cgf.SetHeader("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36")
	var tmpurl bytes.Buffer
	if target.Port == "443" {
		tmpurl = MergeUrl("https://", src.String())
	} else {
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

		//fmt.Println("https:titil update", target.Ip, target.Port, resp.Request.URL)
		dataUpdate := make(map[string]interface{})
		dataUpdate["title"] = title
		dataUpdate["loginurl"] = strings.TrimSuffix(resp.Request.URL.String(), "/")
		models.EditIplistByIp(target.Ip, target.Port, dataUpdate)

		return tmpurl.String()

	} else {
		return ""
	}

}

//根据关键字判断登陆后台
func findLoginWeb() {
	setting := models.GetSettingTitle()
	iplistTitle := models.GetIplistTitle()

	for _, settingKey := range setting {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("findLoginWeb err :", err)
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

// 自动检测指纹
func CheckFinger(HttpRes []string) {
	domains := make(chan string)

	allFinger := models.GetAllFinger()

	var allFingers bytes.Buffer
	allFingers.WriteString(`{"technologies": {`)

	for _, finger := range allFinger {
		allFingers.WriteString(finger.Finger)
		allFingers.WriteString(",")
	}
	Fingers := strings.TrimSuffix(allFingers.String(), ",")
	Fingers = Fingers + " }}"
	//fmt.Println("Fingers:",Fingers)

	if wa, err = alyaze.NewWebAnalyzer(Fingers, nil); err != nil {
		log.Printf("initialization failed: %v", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < thread; i++ {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("[+] find error :", err)
			}
		}()

		wg.Add(1)
		go func() {

			for host := range domains {
				job := alyaze.NewOnlineJob(host, "", nil, crawlCount, searchSubdomain, redirect)
				result, links := wa.Process(job)

				if searchSubdomain {
					for _, v := range links {
						crawlJob := alyaze.NewOnlineJob(v, "", nil, 0, false, redirect)
						wa.Process(crawlJob)
					}
				}

				for _, a := range result.Matches {
					dataUpdate := make(map[string]interface{})
					dataUpdate["cms"] = a.AppName
					models.EditIplistByUrl(result.Host, dataUpdate)
				}

			}

			wg.Done()
		}()
	}
	for _, k := range HttpRes {
		domains <- k
	}

	close(domains)
	wg.Wait()

}

// 测试单个指纹是否编写正常
func CheckFingerSingle(domain, finger string) bool {

	var allFingers bytes.Buffer
	allFingers.WriteString(`{"technologies": {`)

	allFingers.WriteString(finger)
	Fingers := strings.TrimSuffix(allFingers.String(), ",")
	Fingers = Fingers + " }}"
	//fmt.Println("Fingers:",Fingers)

	if wa, err = alyaze.NewWebAnalyzer(Fingers, nil); err != nil {
		log.Printf("initialization failed: %v", err)
		return false
	}

	job := alyaze.NewOnlineJob(domain, "", nil, crawlCount, searchSubdomain, redirect)
	result, _ := wa.Process(job)
	if len(result.Matches) > 0 {
		return true
	}
	return false

}

// 扫描数据库资产在指纹的结果
func CheckFingerOne(id int) {
	var success int

	start := time.Now()
	allFinger := models.GetAllFingerId(id)
	HttpRes := models.GetIplistHttp()

	costTime := time.Since(start)

	data := make(map[string]interface{})
	data["taskid"] = 0
	data["task_name"] = "checkFinger"
	data["task_type"] = "checkFinger"
	data["all_num"] = 0
	data["succes_num"] = 0
	data["run_time"] = ""
	data["error"] = "nil"
	data["status"] = 0
	taskId := models.AddLog(data)

	domains := make(chan string)

	var allFingers bytes.Buffer
	allFingers.WriteString(`{"technologies": {`)

	for _, finger := range allFinger {
		allFingers.WriteString(finger.Finger)
		allFingers.WriteString(",")
	}
	Fingers := strings.TrimSuffix(allFingers.String(), ",")
	Fingers = Fingers + " }}"
	//fmt.Println("Fingers:",Fingers)

	if wa, err = alyaze.NewWebAnalyzer(Fingers, nil); err != nil {
		log.Printf("initialization failed: %v", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < thread; i++ {
		wg.Add(1)
		go func() {

			for host := range domains {
				fmt.Println("now is :", host)
				job := alyaze.NewOnlineJob(host, "", nil, crawlCount, searchSubdomain, redirect)
				result, links := wa.Process(job)

				if searchSubdomain {
					for _, v := range links {
						crawlJob := alyaze.NewOnlineJob(v, "", nil, 0, false, redirect)
						wa.Process(crawlJob)
					}
				}

				for _, a := range result.Matches {
					success++
					dataUpdate := make(map[string]interface{})
					dataUpdate["cms"] = a.AppName
					models.EditIplistByUrl(result.Host, dataUpdate)
				}

			}

			wg.Done()
		}()
	}
	for _, k := range HttpRes {
		if k.Loginurl != "" {
			domains <- k.Loginurl
		}
	}

	close(domains)
	wg.Wait()

	data = make(map[string]interface{})
	data["all_num"] = success
	data["status"] = 1
	data["run_time"] = fmt.Sprintf("%s", costTime)
	models.EditLog(taskId, data)

}

func ScanPocXray(pocName string, targetString []string, thread int) error {

	osType := runtime.GOOS
	nowPath, _ := os.Getwd()

	if _, err := os.Stat("output"); os.IsNotExist(err) {
		_ = os.Mkdir("output", os.ModePerm)
	}
	if _, err := os.Stat(filepath.Join("output", "xrayinput")); os.IsNotExist(err) {
		_ = os.Mkdir(filepath.Join("output", "xrayinput"), os.ModePerm)
	}
	if _, err := os.Stat(filepath.Join("output", "xrayout")); os.IsNotExist(err) {
		_ = os.Mkdir(filepath.Join("output", "xrayout"), os.ModePerm)
	}

	xrayOuputPaht := filepath.Join(nowPath, "output", "xrayinput")
	fileName := common.GetRandomString(12)

	fileName = filepath.Join(xrayOuputPaht, fileName+".txt")
	common.WirteFileAppend(fileName, targetString)

	// 输出文件路径
	outJson := time.Now().Format("2006-01-02 15:04:05") + ".json"
	xrayOuputputh := filepath.Join(nowPath, "output", "xrayout", outJson)

	cmd := &exec.Cmd{}

	if osType == "darwin" {
		// 要运行的程序名称，比如 ksubdomainMac
		runName := filepath.Join(nowPath, "pkg", "third", "xray_darwin_amd64")
		//./xray_darwin_amd64 webscan --url http://127.0.0.1:19001  --html-output 0102.html
		// --webhook-output http://127.0.0.1:5000/webhook
		if pocName != "" {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--poc", pocName, "--json-output", xrayOuputputh)
		} else {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--json-output", xrayOuputputh)
		}

	} else if osType == "linux" {
		runName := filepath.Join(nowPath, "pkg", "third", "xray_linux_amd64")
		if pocName != "" {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--poc", pocName, "--json-output", xrayOuputputh)
		} else {
			cmd = exec.Command(runName, "webscan", "--url-file", fileName, "--json-output", xrayOuputputh)
		}
	} else {
		fmt.Println("[+] 运行xray命令错误", osType)
	}

	fmt.Println("\n\nxray_darwin_amd64 webscan runcmd:", cmd)

	cmdOutput, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("runcmd err : " + fmt.Sprint(err) + ":" + string(cmdOutput))
	}

	//if common.PathExist(fileName) {
	//	os.Remove(fileName)
	//}
	//
	if common.PathExist(xrayOuputputh) {
		XrayRes(xrayOuputputh)
		//os.Remove(xrayOuputputh)
	}

	return nil

}

func XrayRes(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("xray into db err :", err)
			}
		}()
		var xrayResult XrayResAll

		lines, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return
			}
			return
		}
		decoder := json.NewDecoder(strings.NewReader(lines))
		err = decoder.Decode(&xrayResult)
		if err != nil {
			fmt.Println("Decoder failed : ", err.Error())
		} else {
			snapshot := common.SliceSToString(xrayResult.Detail.Snapshot)
			hash := common.GetMD5Hash(xrayResult.Target.URL+xrayResult.Plugin)
			lock.Lock()
			data := make(map[string]interface{})
			data["url"] = xrayResult.Target.URL
			data["poc"] = xrayResult.Plugin
			data["hash"] = hash
			data["snapshot"] = snapshot
			models.AddXrayres(data)
			lock.Unlock()
		}
	}
}
