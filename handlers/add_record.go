package handlers

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	jd "github.com/josephburnett/jd/lib"

	"github.com/bubble-diff/bubblereplay/app"
	"github.com/bubble-diff/bubblereplay/levenshtein"
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
	go addRecordProcess(c, record)

	c.JSON(200, nil)
}

func addRecordProcess(ctx context.Context, record *models.Record) {
	var err error
	// todo: [Advanced] 每次add record都要查task配置太不划算，应该将这个信息缓存起来，
	//  尝试以TaskID为最小单位建立goroutine。
	task, err := app.GetTaskDetail(record.TaskID)
	if err != nil {
		log.Println(err)
		return
	}

	// 读取旧请求
	oldReq, err := http.ReadRequest(bufio.NewReader(bytes.NewReader(record.OldReq)))
	if err != nil {
		log.Println(err)
		return
	}

	// 过滤判断
	if task.FilterConfig.Drop(ctx, oldReq) {
		log.Printf("[addRecordProcess] drop packet=%+v", oldReq)
		return
	}

	// 更新record总数
	err = app.AddTotalRecord(ctx, task.ID)
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
	var metadata []jd.Metadata
	if task.AdvanceConfig.IsIgnoreArraySequence {
		metadata = append(metadata, jd.SET)
	}
	oldNode, _ := jd.ReadJsonString(string(record.OldResp))
	newNode, _ := jd.ReadJsonString(string(record.NewResp))
	record.Diff = oldNode.Diff(newNode, metadata...).Render()

	// 计算差异百分比
	if len(record.Diff) != 0 {
		record.DiffRate = levenshtein.Compute(string(record.OldResp), string(record.NewResp))
	}

	// 上传record至cos
	cosKey, err := app.UploadRecord(ctx, record)
	if err != nil {
		log.Println(err)
		return
	}

	err = app.AppendRecordMeta(ctx, record.TaskID, cosKey, oldReq.URL.Path, record.DiffRate)
	if err != nil {
		log.Println(err)
		return
	}

	err = app.AddSuccessRecord(ctx, task.ID)
	if err != nil {
		log.Println(err)
		return
	}
}
