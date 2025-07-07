package ctl

import (
	"micro-memorandum/pkg/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status int         `json:"status, omitempty"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

func RespSucess(ctx *gin.Context, data interface{}, code ...int) *Response {
	status := e.Success
	if code != nil {
		status = code[0]
	}

	if data == nil {
		data = "操作成功"
	}

	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
	}
}

func RespError(ctx *gin.Context, err error, data string, code ...int) *Response {
	status := e.Error
	if code != nil {
		status = code[0]
	}

	return &Response{
		Status: status,
		Data:   data,
		Msg:    e.GetMsg(status),
		Error:  err.Error(),
	}
}
