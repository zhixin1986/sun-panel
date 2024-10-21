package main

import (
	"log"
	"sun-panel/global"
	"sun-panel/initialize"
	"sun-panel/router"
)

func main() {
	err := initialize.InitApp()
	if err != nil {
		log.Println("初始化错误:", err.Error())
		panic(err)
	}
	httpPort := global.Config.GetValueStringOrDefault("base", "http_port")
	httpsPort := global.Config.GetValueStringOrDefault("base", "https_port")
	sslEnable := global.Config.GetValueStringOrDefault("base", "ssl_enable")
	if sslEnable == "true" {
		if err := router.InitRouters(":"+httpsPort, true); err != nil {
			log.Println("InitRouters 错误:", err.Error())
			panic(err)
		}
	} else {
		if err := router.InitRouters(":"+httpPort, false); err != nil {
			log.Println("InitRouters 错误:", err.Error())
			panic(err)
		}
	}

}
