package log

import (
	"github.com/lucky-cheerful-man/phoenix_gateway/src/config"
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	logger := &lumberjack.Logger{
		Filename:   config.GetGlobalConfig().AppSetting.LogFileName,
		MaxSize:    config.GetGlobalConfig().AppSetting.LogMaxSize,    // 日志文件大小，单位是 MB
		MaxBackups: config.GetGlobalConfig().AppSetting.LogMaxBackups, // 最大过期日志保留个数
		MaxAge:     config.GetGlobalConfig().AppSetting.LogMaxAgeDay,  // 保留过期文件最大时间，单位 天
		Compress:   config.GetGlobalConfig().AppSetting.LogCompress,   // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
	}

	log.SetOutput(logger)
	log.SetLevel(getLevel(config.GetGlobalConfig().AppSetting.LogLevel))
	log.SetReportCaller(true)
	log.SetFormatter(&logrus.JSONFormatter{})
}

// 获取本次启用的日志级别
func getLevel(level string) logrus.Level {
	switch level {
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	case "FATAL":
		return logrus.FatalLevel
	}

	return logrus.InfoLevel
}

// Infof 封装一层info日志
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf 封装一层warn日志
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf 封装一层error日志
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf 封装一层Fatal日志
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
