package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/constant"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/rpc"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/util"
	"net/http"
)

type registerInfo struct {
	Name     string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MinSize(6); MaxSize(50)"`
}

// Register 注册
func Register(c *gin.Context) {
	appG := util.Gin{C: c}
	id, ok := c.Get("requestId")
	if !ok {
		log.Errorf("get requestId failed")
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}
	requestID := id.(string)
	valid := validation.Validation{}

	name := c.PostForm("name")
	password := c.PostForm("password")

	a := registerInfo{Name: name, Password: password}
	ok, _ = valid.Valid(&a)
	if !ok {
		util.MarkErrors(requestID, valid.Errors)
		appG.Response(http.StatusBadRequest, constant.InvalidParams, nil)
		return
	}

	err := rpc.Register(requestID, name, password)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.ErrorRegisterFailed, nil)
		return
	}

	appG.Response(http.StatusOK, constant.Success, map[string]string{})
}
