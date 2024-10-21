package router

import (
	"net/http"
	"sun-panel/global"

	// "sun-panel/router/admin"
	"sun-panel/router/openness"
	"sun-panel/router/panel"
	"sun-panel/router/system"

	"github.com/gin-gonic/gin"
)

// 初始化总路由
func InitRouters(addr string, ssl bool) error {
	router := gin.Default()
	rootRouter := router.Group("/")
	routerGroup := rootRouter.Group("api")

	// 接口
	system.Init(routerGroup)
	panel.Init(routerGroup)
	openness.Init(routerGroup)

	// WEB文件服务
	{
		webPath := "./web"
		router.StaticFile("/", webPath+"/index.html")
		router.Static("/assets", webPath+"/assets")
		router.Static("/custom", webPath+"/custom")
		router.StaticFile("/favicon.ico", webPath+"/favicon.ico")
		router.StaticFile("/favicon.svg", webPath+"/favicon.svg")
	}

	// 上传的文件
	sourcePath := global.Config.GetValueString("base", "source_path")
	router.Static(sourcePath[1:], sourcePath)
	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	if ssl {
		global.Logger.Info("Sun-Panel is Started.  Listening and serving HTTPS on ", addr)
		return srv.ListenAndServeTLS("conf/ssl.cert", "conf/ssl.key")
	} else {
		global.Logger.Info("Sun-Panel is Started.  Listening and serving HTTP on ", addr)
		return router.Run(addr)
	}

}
