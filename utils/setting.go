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
	file, err := ini.Load("config/my.ini")
	if err != nil {
		fmt.Println("配置文件读取出错", err)
	}
	loadServer(file)
	loadDb(file)
}

// 获取服务器设置
func loadServer(file *ini.File) {
	// MustString 默认设置
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
}

// 获取数据库配置
func loadDb(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("mysql")
	DbHost = file.Section("database").Key("DbHost").MustString("81.70.27.234")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("root")
	DbPassword = file.Section("database").Key("DbPassWord").MustString("010825lwj")
	DbName = file.Section("database").Key("DbName").MustString("ginBlog")
}
