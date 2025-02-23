package util

import (
	"strings"

	"github.com/sunshine-walker-93/phoenix_gateway/src/config"
)

// CheckImageSize 检查图片大小是否合法
func CheckImageSize(size int) bool {
	return size <= config.GetGlobalConfig().AppSetting.ImageMaxSize
}

// CheckImageExt 检查扩展是否合法
func CheckImageExt(fileName string) bool {
	ext := GetExt(fileName)
	for _, allowExt := range config.GetGlobalConfig().AppSetting.ImageAllowExt {
		if strings.EqualFold(allowExt, ext) {
			return true
		}
	}
	return false
}
