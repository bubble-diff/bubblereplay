package handlers

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	jd "github.com/josephburnett/jd/lib"

	"github.com/bubble-diff/bubblereplay/app"
	"github.com/bubble-diff/bubblereplay/pb"
)

func AddRecord(c *gin.Context) {
	var err error
	var req pb.AddRecordReq
	var resp pb.AddRecordResp

	err = binding.ProtoBuf.Bind(c.Request, &req)
	if err != nil {
		resp.Code = -1
		resp.Msg = err.Error()
		c.ProtoBuf(200, &resp)
		return
	}

	// 后台异步重放，忽略操作失败，因为丢失一些"record"是可容忍的。
	// todo: [Advanced] 不必每个record都建立goroutine，
	//  尝试以TaskID为最小单位建立goroutine。
	go addRecordProcess(req.Record)

	resp.Msg = "OK"
	c.ProtoBuf(200, &resp)
}

func addRecordProcess(record *pb.Record) {
	// todo: [Advanced] 每次add record都要查task配置太不划算，应该将这个信息缓存起来，
	//  尝试以TaskID为最小单位建立goroutine。
	task, err := app.GetTaskDetail(record.TaskId)

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

	rawNewResp, err := io.ReadAll(newResp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	record.NewResp = rawNewResp

	// 差异比对新旧响应
	// todo: 根据diff任务配置进行高级处理
	oldNode, _ := jd.ReadJsonString(string(record.OldResp))
	newNode, _ := jd.ReadJsonString(string(record.NewResp))
	result := oldNode.Diff(newNode)
	log.Println(result.Render())

	// todo: [Advanced] 将record存储在腾讯cos中
	//  目前为了方便存在redis里先。

}
