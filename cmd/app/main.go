package main

import (
	v1 "GoToDoList/internal/api/v1"
	"GoToDoList/internal/global"
	"GoToDoList/internal/initialize"
)

func main() {
	initialize.SetUpViper()
	initialize.SetupLogger()
	initialize.SetupDatabase()

	router := v1.Router()

	global.Logger.Info("服务正在启动...端口在8080...")
	if err := router.Run("localhost:8080"); err != nil {
		global.Logger.Error("服务启动失败" + err.Error())
	}
}
