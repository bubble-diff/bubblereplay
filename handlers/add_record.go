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

	go addRecordProcess(req.Record)

	resp.Msg = "OK"
	c.ProtoBuf(200, &resp)
}

// addRecordProcess 都是耗时操作，需要异步完成(goroutine)，忽略操作失败，因为丢失record是可容忍的。
func addRecordProcess(record *pb.Record) {
	// todo: 第一步，重放req并diff，最后日志打印diff结果

	// read raw_req
	oldReq, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(record.OldReq)))
	if err != nil {
		log.Println(err)
		return
	}
	// build new req
	_url := oldReq.URL
	_url.Scheme = "http"
	// todo: new req host get from task config
	_url.Host = "127.0.0.1:8081"
	newReq, err := http.NewRequest(oldReq.Method, _url.String(), oldReq.Body)
	for key, values := range oldReq.Header {
		for _, value := range values {
			newReq.Header.Add(key, value)
		}
	}

	// get new resp
	newResp, err := http.DefaultClient.Do(newReq)
	if err != nil {
		log.Println(err)
		return
	}
	// todo: if newResp is nil
	if newResp == nil {
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

	// diff old/new resp
	// todo: 根据diff任务配置进行高级处理
	oldNode, _ := jd.ReadJsonString(string(record.OldResp))
	newNode, _ := jd.ReadJsonString(string(record.NewResp))
	result := oldNode.Diff(newNode)
	log.Println(result.Render())

	// todo: 第二步，将record存储在腾讯cos中
}
