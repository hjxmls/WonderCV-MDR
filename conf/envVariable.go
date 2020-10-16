package conf

import (
	"github.com/gogf/gf/os/genv"
)

var (
	//redis
	REDIS_URL string
)

//初始化从环境变量中读取的全局变量
func InitEnv() {
	//redis
	REDIS_URL = genv.Get("REDIS_URL", "https://localhost:6379")

}
