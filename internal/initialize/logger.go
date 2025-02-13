package initialize

import (
	"layerg-poi-da/global"
	"layerg-poi-da/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
