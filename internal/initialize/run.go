package initialize

import (
	"layerg-poi-da/global"
)

func Run() {
	InitLogger()
	InitDB()
	InitRedis()

	r := InitRouter()

	r.Run(":" + global.Config.Server.Port)
}
