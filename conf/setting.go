package conf

import (
	"github.com/go-ini/ini"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

var (
	Cfg        *ini.File
	SERVER_ENV string
	RunMode    string

	//Server
	SERVER_HOST          string
	SERVER_PORT          string
	SERVER_READ_TIMEOUT  time.Duration
	SERVER_WRITE_TIMEOUT time.Duration
	SERVER_DOMAIN_NAME   string
)

func InitConf() {
	file, _ := os.Getwd()
	var err error
	var ok bool
	SERVER_ENV, ok = os.LookupEnv("SERVER_ENV")
	if !ok {
		SERVER_ENV = "dev"
	}
	path := file + "/conf/conf.ini"
	Cfg, err = ini.Load(path)
	if err != nil {
		logrus.Errorf("Fail to parse '%v' ", path)
	}
	LoadBase()
	LoadServer()
}

func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").String()
}

func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v\n", err)
	}
	SERVER_HOST = sec.Key("SERVER_HOST").MustString("127.0.0.1")
	SERVER_PORT = sec.Key("SERVER_PORT").MustString("8080")
	SERVER_READ_TIMEOUT = time.Duration(sec.Key("SERVER_READ_TIMEOUT").MustInt(10)) * time.Second
	SERVER_WRITE_TIMEOUT = time.Duration(sec.Key("SERVER_WRITE_TIMEOUT").MustInt(10)) * time.Second
	SERVER_DOMAIN_NAME = sec.Key("SERVER_DOMAIN_NAME").MustString("8080")
}
