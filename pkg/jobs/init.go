package jobs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"linglong/models"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func InitJobs() {
	list := models.TaskInitList(1, 1000000, 1)

	for _, task := range list {
		job, err := NewJobFromTask(task)
		if err != nil {
			continue
		}
		AddJob(task.CronSpec, job)
	}
}

func HttpGet(timeout time.Duration, task *models.Task) (string, string, error, bool) {
	arge := task.Arge + "&taskid=" + strconv.Itoa(task.Id)
	return "", "", nil, requesPostwww(task.Command, arge)
}

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func requestGet(url string, timeout time.Duration) (string, string, error, bool) {

	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	resp, err := client.Get(url)

	var isTimeout bool
	if err != nil {
		isTimeout = true
	} else {
		isTimeout = false
	}
	if resp == nil {
		return "", "", err, false
	}
	// status_code:= resp.StatusCode //获取返回状态码
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)

	return string(body), "", err2, isTimeout

}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json   application/x-www-form-unlencoded
// content：     请求放回的内容
func requestPost(url string, data interface{}, timeout time.Duration) (string, string, error, bool) {
	contentType := "application/x-www-form-unlencoded"
	// 超时时间：5秒
	client := &http.Client{Timeout: time.Duration(timeout) * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	var isTimeout bool
	if err != nil {
		isTimeout = true
	} else {
		isTimeout = false
	}
	defer resp.Body.Close()

	result, err2 := ioutil.ReadAll(resp.Body)
	return string(result), err2.Error(), err2, isTimeout
}

func requesPostwww(url, data string) bool {
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		strings.NewReader(data))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
	return true
}
