package v1

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/constant"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/log"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/rpc"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/util"
	"io"
	"net/http"
)

// GetProfile 查询用户基本信息
func GetProfile(c *gin.Context) {
	appG := util.Gin{C: c}
	cname, ok := c.Get("name")
	if !ok {
		appG.Response(http.StatusBadRequest, constant.InvalidParams, nil)
		return
	}
	name := cname.(string)

	id, ok := c.Get("requestId")
	if !ok {
		log.Errorf("get requestId failed")
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}
	requestID := id.(string)

	res, err := rpc.GetProfile(requestID, name)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}

	appG.Response(http.StatusOK, constant.Success, map[string]string{
		"nickname": res.Nickname,
		"imageID":  res.ImageID,
	})
}

// GetHeadImage 查询用户图片详情
func GetHeadImage(c *gin.Context) {
	appG := util.Gin{C: c}
	id, ok := c.Get("requestId")
	if !ok {
		log.Errorf("get requestId failed")
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}
	requestID := id.(string)

	imageID := c.Query("imageID")
	if len(imageID) == 0 {
		log.Warnf("request:%s imageID length invalid", requestID)
		appG.Response(http.StatusBadRequest, constant.InvalidParams, nil)
		return
	}

	res, err := rpc.GetHeadImage(requestID, imageID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}

	buf := bytes.NewBuffer(res)
	size := buf.Len()
	count, err := io.Copy(c.Writer, buf)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}
	if int(count) != size {
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", imageID))
	c.Header("Content-Type", "application/text/plain")
	c.Header("Accept-Length", fmt.Sprintf("%d", size))
}

// EditProfile 编辑用户的属性信息
//nolint:funlen
func EditProfile(c *gin.Context) {
	appG := util.Gin{C: c}
	cname, ok := c.Get("name")
	if !ok {
		appG.Response(http.StatusBadRequest, constant.InvalidParams, nil)
		return
	}
	name := cname.(string)

	id, ok := c.Get("requestId")
	if !ok {
		log.Errorf("get requestId failed")
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}
	requestID := id.(string)

	nickname := c.PostForm("nickname")

	buf := bytes.NewBuffer(nil)
	file, err := c.FormFile("image")
	if err == nil {
		if !util.CheckImageExt(file.Filename) {
			log.Infof("%s invalid image type %s", requestID, file.Filename)
			appG.Response(http.StatusBadRequest, constant.ErrorUploadCheckImageExt, nil)
			return
		}

		if !util.CheckImageSize(int(file.Size)) {
			log.Infof("%s invalid image size %d", requestID, file.Size)
			appG.Response(http.StatusBadRequest, constant.ErrorUploadCheckImageFormat, nil)
			return
		}

		src, err := file.Open()
		if err != nil {
			log.Warnf("%s open file failed:%s", requestID, err)
			appG.Response(http.StatusInternalServerError, constant.Error, nil)
			return
		}
		defer func() {
			err = src.Close()
			if err != nil {
				log.Warnf("%s close file failed,err:%s", requestID, err)
			}
		}()

		// 读取file的数据存入buf中
		_, err = io.Copy(buf, src)
		if err != nil {
			log.Warnf("%s copy file failed:%s", requestID, err)
			appG.Response(http.StatusInternalServerError, constant.Error, nil)
			return
		}
	} else if err.Error() != "http: no such file" && err.Error() != "request Content-Type isn't multipart/form-data" {
		log.Warnf("%s get file failed:%s", requestID, err)
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}

	err = rpc.EditProfile(requestID, name, nickname, buf.Bytes())
	if err != nil {
		appG.Response(http.StatusInternalServerError, constant.Error, nil)
		return
	}

	appG.Response(http.StatusOK, constant.Success, map[string]string{})
}
