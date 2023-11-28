package Dao

import "ToDoList/model"

func migration() {
	DB.Set("gorm:table_options", "charset=utf8mb4").
		AutoMigrate(&model.UserDao{}).
		AutoMigrate(&model.TaskDao{})
	DB.Model(&model.TaskDao{}).AddForeignKey("uid", "user_dao(id)", "CASCADE", "CASCADE")
	//CASCADE:在父表上update/delete记录时，同步update/delete子表的匹配记录
}
