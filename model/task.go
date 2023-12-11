package model

import (
	"github.com/jinzhu/gorm"
)

type TaskDao struct {
	gorm.Model
	User      UserDao `gorm:"ForeignKey:Uid"`
	UserName  string  `gorm:"not null"`
	Tid       int     `gorm:"not null"`
	Uid       uint    `gorm:"not null"`
	Title     string  `gorm:"index;not null"`
	View      int     `gorm:"default:'0'"`
	Status    int     `gorm:"default:'0'"` //0未完成 1已完成
	Content   string  `gorm:"type:longtext"`
	StartTime string
	EndTime   string
}

type Task struct {
	UserName  string `json:"user_name" form:"user_name"`
	Uid       uint   `gorm:"not null"`
	Tid       int    `json:"tid" form:"tid"`
	Title     string `json:"title" form:"title"`
	View      int    `json:"view" form:"view"`
	Status    int    `json:"status" form:"status"` //0未完成 1已完成
	Content   string `json:"content" form:"content"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
}

func BuildTask(task TaskDao) Task {
	return Task{
		UserName:  task.UserName,
		Uid:       task.Uid,
		Tid:       task.Tid,
		Title:     task.Title,
		Status:    task.Status,
		View:      task.View,
		Content:   task.Content,
		StartTime: task.StartTime,
		EndTime:   task.EndTime,
	}
}
