package main

import (
	"GolangStudy/Practice/shop/user-web/global"
	"GolangStudy/Practice/shop/user-web/initialize"
	"fmt"

	"go.uber.org/zap"
)

func main() {
	//1.初始化logger
	initialize.InitLogger()
	initialize.InitConfig()
	//1.初始化routers
	Router := initialize.Routers()
	zap.S().Infof("启动服务器,端口:%d", global.ServerConfig.Port)
	Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port))
}
