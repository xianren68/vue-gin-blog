package utils

import (
	"fmt"

	"github.com/go-ini/ini"
)

// 定义一些变量
var (
	AppMode    string
	HttpPort   string
	Db         string
	DbPort     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
)

func init() {
	file, err := ini.Load("../config/my.ini")
	if err != nil {
		fmt.Println("配置文件读取出错", err)
	}
}
