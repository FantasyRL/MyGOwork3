package Dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

// 该文件用于进行数据库的连接

var DB *gorm.DB

// 初始化数据库
func Database(connstring string) {
	db, err := gorm.Open("mysql", connstring)
	if err != nil {
		panic("数据库连接错误")
	}
	fmt.Println("mysql connect access")
	db.LogMode(true)             //日志
	db.SingularTable(true)       //表名末尾不加s（mysql会在表名末尾+s）
	db.DB().SetMaxIdleConns(20)  //设置连接池
	db.DB().SetMaxOpenConns(100) //最大连接数
	db.DB().SetConnMaxLifetime(time.Second * 30)
	DB = db
	migration()
}
