package handlers

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	jd "github.com/josephburnett/jd/lib"

	"github.com/bubble-diff/bubblereplay/app"
	"github.com/bubble-diff/bubblereplay/models"
)

func AddRecord(c *gin.Context) {
	var err error
	record := new(models.Record)
	err = c.BindJSON(record)
	if err != nil {
		log.Printf("[AddRecord] unmarshal json failed, %s", err)
		c.JSON(200, nil)
		return
	}

	// 后台异步重放，忽略操作失败，因为丢失一些"record"是可容忍的。
	// todo: [Advanced] 不必每个record都建立goroutine，
	//  尝试以TaskID为最小单位建立goroutine。
	go addRecordProcess(record)

	c.JSON(200, nil)
}

func addRecordProcess(record *models.Record) {
	var err error
	// todo: [Advanced] 每次add record都要查task配置太不划算，应该将这个信息缓存起来，
	//  尝试以TaskID为最小单位建立goroutine。
	task, err := app.GetTaskDetail(record.TaskID)

	// 读取旧请求
	oldReq, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(record.OldReq)))
	if err != nil {
		log.Println(err)
		return
	}
	// 构建新请求
	_url := oldReq.URL
	_url.Scheme = "http"
	_url.Host = task.TrafficConfig.Addr
	newReq, err := http.NewRequest(oldReq.Method, _url.String(), oldReq.Body)
	for key, values := range oldReq.Header {
		for _, value := range values {
			newReq.Header.Add(key, value)
		}
	}
	// 获得新响应
	newResp, err := http.DefaultClient.Do(newReq)
	if err != nil {
		log.Println(err)
		return
	}

	// todo: if Content-Encoding specified, set an appropriate reader,
	//  like gzip, deflate...

	// todo: use an appropriate charset to read body

	record.NewResp, err = io.ReadAll(newResp.Body)
	if err != nil {
		log.Println(err)
		return
	}

	// 差异比对新旧响应
	// todo: 根据diff任务配置进行高级处理
	oldNode, _ := jd.ReadJsonString(string(record.OldResp))
	newNode, _ := jd.ReadJsonString(string(record.NewResp))
	record.Diff = oldNode.Diff(newNode).Render()
	log.Println(record.Diff)

	// todo: [Advanced] 将record存储在腾讯cos中
	//  目前为了方便存在redis里先。

}
