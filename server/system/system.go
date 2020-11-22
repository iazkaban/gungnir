package system

import (
	"gungnir/models"
	"runtime"
)

var (
	info    *models.SystemInfo
	Version string
)

func init() {
	info = &models.SystemInfo{
		GoVersion: runtime.Version(),
		Version:   Version,
	}
}

func GetSystemInfo() *models.SystemInfo {
	return info
}
