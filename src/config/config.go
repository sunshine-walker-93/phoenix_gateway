package config

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret       string
	DeadlineSecond  int
	RuntimeRootPath string
	TokenExpireHour int
	MaxHeaderBytes  int

	ImageMaxSize  int
	ImageAllowExt []string

	LogLevel      string // 日志文件展示级别
	LogFileName   string // 日志文件存放路径与名称
	LogMaxSize    int    // 日志文件大小，单位是 MB
	LogMaxBackups int    // 最大过期日志保留个数
	LogMaxAgeDay  int    // 保留过期文件最大时间，单位 天
	LogCompress   bool   // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type DaoServer struct {
	GrpcPort int
}

type GlobalConfig struct {
	cfg              *ini.File
	AppSetting       App
	ServerSetting    Server
	DaoServerSetting DaoServer
}

var globalConfig *GlobalConfig

func init() {
	var err error
	globalConfig = new(GlobalConfig)
	globalConfig.cfg, err = ini.Load("localconf/app.ini")
	if err != nil {
		fmt.Printf("config.Init, fail to parse app.ini: %s", err)
		panic("read config file failed")
	}

	mapTo("app", globalConfig.AppSetting)
	mapTo("server", globalConfig.ServerSetting)
	mapTo("dao-server", globalConfig.DaoServerSetting)

	globalConfig.ServerSetting.ReadTimeout *= time.Second
	globalConfig.ServerSetting.WriteTimeout *= time.Second
}

// mapTo map section
func mapTo(section string, v interface{}) {
	err := globalConfig.cfg.Section(section).MapTo(v)
	if err != nil {
		fmt.Printf("Cfg.MapTo %s err: %s", section, err)
		panic("get config failed")
	}
}

func GetGlobalConfig() *GlobalConfig {
	return globalConfig
}
