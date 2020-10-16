package main

import (
	"WonderCV-MDR/conf"
	"WonderCV-MDR/initialize"
	"WonderCV-MDR/middlewares"
	"WonderCV-MDR/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	//初始化配置
	initialize.Initialize()
}

func main() {
	//routers.WsConnManager.Close()
	////断续器，定时刷新accesstoken
	//var ticker *time.Ticker = time.NewTicker(600 * time.Second)
	//
	//defer ticker.Stop()
	r := gin.Default()
	//引入middleware
	r.Use(middlewares.Cors(), middlewares.Timeout(conf.SERVER_READ_TIMEOUT))
	//r.Use(auth.Auth)

	//routers
	router := r.Group("/mdr")
	v1 := router.Group("/v1")

	v1.POST("/resumes", routers.GenResumes)

	//runserver
	_ = r.Run(":" + conf.SERVER_PORT)
}
