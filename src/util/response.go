package util

import (
	"github.com/gin-gonic/gin"

	"github.com/lucky-cheerful-man/phoenix_gateway/src/constant"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int32       `json:"constant"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response config gin.JSON
func (g *Gin) Response(httpCode int, errCode constant.ErrorStruct, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: errCode.Code,
		Msg:  errCode.Msg,
		Data: data,
	})
}

// ProfileInfo 用户的属性信息
type ProfileInfo struct {
	Nickname string
	ImageID  string
}
