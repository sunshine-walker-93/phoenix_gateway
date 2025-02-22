package routers

import (
	"github.com/gin-gonic/gin" //nolint:goimports
	"github.com/lucky-cheerful-man/phoenix_gateway/src/config"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/middleware/jwt"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/middleware/request"
	api2 "github.com/lucky-cheerful-man/phoenix_gateway/src/routers/api"
	"github.com/lucky-cheerful-man/phoenix_gateway/src/routers/api/v1"
)

func init() {
	gin.SetMode(config.GetGlobalConfig().ServerSetting.RunMode)
}

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(request.GenRequestID())
	r.POST("/register", api2.Register)
	r.POST("/login", api2.GetAuth)

	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		// 获取用户基本信息
		apiV1.GET("/profile", v1.GetProfile)
		// 获取头像信息
		apiV1.GET("/profile/image", v1.GetHeadImage)
		// 编辑用户信息
		apiV1.POST("/profile", v1.EditProfile)
	}

	return r
}
