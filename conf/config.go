package conf

import (
	"ToDoList/Dao"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/ini.v1"
	"strings"
)

// go get gopkg.in/ini.v1
var (
	AppMode     string
	HttpPort    string
	RedisAddr   string
	RedisPw     string
	RedisDbName string
	Db          string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassWord  string
	DbName      string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误")
	}
	LoadServer(file)
	LoadMysql(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	Dao.Database(path)
}

// LoadServer 加载服务器配置
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String() //选中ini中的service模块的AppMode
	HttpPort = file.Section("service").Key("HttpPort").String()
}

// LoadMysql 加载数据库配置
func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}
