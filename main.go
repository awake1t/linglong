package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"linglong/global"
	"linglong/models"
	"linglong/pkg/jobs"
	"linglong/pkg/setting"
	"linglong/routers"
	v1 "linglong/routers/api/v1"
	"log"
	"strings"
	"syscall"
	"time"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	jobs.InitJobs()

}
func main() {

	endless.DefaultReadTimeOut = global.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = global.ServerSetting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20

	endPoint := fmt.Sprintf(":%d", global.ServerSetting.HttpPort)

	server := endless.NewServer(endPoint, routers.InitRouter())
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}

	//启动一个无限循环扫描的masscan+nmap端口识别
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("masscan 运行错误:", err)
			}
		}()
		i := 1
		for {
			start := time.Now()
			v1.InitMasscan()
			end := time.Now()
			log.Printf("masscan完成第 %d 轮扫描 耗时: %s\n", i, end.Sub(start))
			i++
			time.Sleep(10 * time.Second)
		}
	}()

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}

func setupSetting() error {
	s, err := setting.NewSetting(strings.Split("configs/", ",")...)
	if err != nil {
		return err
	}
	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Masscan", &global.MasscanSetting)
	if err != nil {
		return err
	}
	global.AppSetting.DefaultContextTimeout *= time.Second
	//global.JWTSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second

	models.Setup()

	return nil
}
