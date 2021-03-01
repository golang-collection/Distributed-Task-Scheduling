package app

import (
	"Distributed-Task-Scheduling/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
* @Author: super
* @Date: 2021-02-28 13:45
* @Description:
**/

type Response struct {
	Ctx *gin.Context
}

type Meta struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

type Pager struct {
	// 页码
	Page int `json:"page"`
	// 每页数量
	PageSize int `json:"page_size"`
	// 总行数
	TotalRows int `json:"total_rows"`
}

func NewResponse(ctx *gin.Context) *Response {
	return &Response{
		Ctx: ctx,
	}
}

func (r *Response) ToResponse(data interface{}, msg string, status int) {
	if data == nil {
		data = gin.H{}
	} else {
		data = gin.H{
			"data": data,
			"meta": Meta{
				Msg:    msg,
				Status: status,
			},
		}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"data": gin.H{},
		"meta": Meta{
			Msg:    err.Msg(),
			Status: err.Code(),
		},
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}

	r.Ctx.JSON(err.StatusCode(), response)
}
