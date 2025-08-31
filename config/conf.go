package config

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"io"
	"log"
	"os"
)

var (
	ConfPath string
	BeLog    bool
)

func Init() error {
	flag.StringVar(&ConfPath, "conf", "./config/conf.yaml", "配置文件路径")
	flag.BoolVar(&BeLog, "log", false, "是否记录日志在文件")
	flag.Parse()
	viper.SetConfigFile(ConfPath)

	if BeLog {
		setGinLog()
	}

	viper.SetDefault("mysql.dsn", "root:123456@tcp(127.0.0.1:3306)/EventGlide?charset=utf8mb4&parseTime=True&loc=Local")
	viper.SetDefault("mysql.maxIdleConns", 10)
	viper.SetDefault("mysql.maxOpenConns", 100)
	viper.SetDefault("redis.addr", "127.0.0.1:6379")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return nil
}

func setGinLog() {
	os.MkdirAll("./log", 0755)

	f, err := os.OpenFile("./log/gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("无法创建日志文件: %v", err)
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
