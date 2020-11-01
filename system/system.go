package system

import (
	"gungnir/models"
	"runtime"
)

var info *models.SystemInfo

func init() {
	info = &models.SystemInfo{
		GoVersion: runtime.Version(),
		Version:   "1.0.0",
		OS:        runtime.GOOS,
		OSVersion: runtime.GOOS,
	}
}

func GetSystemInfo() *models.SystemInfo {
	return info
}
