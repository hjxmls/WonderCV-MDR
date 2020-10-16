package initialize

import (
	"WonderCV-MDR/conf"
	"WonderCV-MDR/pkg/redis"
	"fmt"
)

func Initialize() {
	//初始化配置
	conf.InitConf()
	//初始化环境变量
	conf.InitEnv()
	//初始化log
	////初始化websocket conn管理
	//connManager.InitWsConnManager()
	//初始化文件读取
	fmt.Printf("SERVER_ENV : %v\n", conf.SERVER_ENV)
	//初始化Redis
	redis.InitRedis()
}
