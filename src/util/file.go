package util

import (
	"path"
)

// GetExt get the file ext
func GetExt(fileName string) string {
	return path.Ext(fileName)
}
