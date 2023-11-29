package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserDao struct {
	gorm.Model            //model提供了ID
	UserName       string `gorm:"unique"`
	PasswordDigest string //存储加密后的密码
	TidCount       int
}

type User struct {
	ID       uint   `json:"id" form:"id" example:"1"`                    //用户ID
	UserName string `json:"user_name" form:"user_name" example:"FanOne"` //用户名
	CreateAt string `json:"create_at" form:"create_at"`                  //创建
}

func BuildUser(user UserDao) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Format(time.RFC3339),
	}
}
