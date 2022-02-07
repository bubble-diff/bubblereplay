package handlers

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

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

	log.Printf("%+v", req.Record)

	resp.Msg = "OK"
	c.ProtoBuf(200, &resp)
}
