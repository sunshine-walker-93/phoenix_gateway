package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	nanoid "github.com/matoous/go-nanoid"
	"github.com/sunshine-walker-93/phoenix_gateway/src/constant"
	"github.com/sunshine-walker-93/phoenix_gateway/src/log"
)

// GenRequestID 生成全局请求id
func GenRequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID, err := nanoid.Nanoid()
		if err != nil {
			log.Warnf("gen nanoid failed:%s", err)

			var data interface{}
			c.JSON(http.StatusInternalServerError, gin.H{
				"constant": constant.Error.Code,
				"msg":      constant.Error.Msg,
				"data":     data,
			})

			c.Abort()
			return
		}
		c.Set("requestId", requestID)
		c.Next()
	}
}
