package AdminService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
)

func (service *AddAdmin) AddAdmin(id string) model.Response {
	var user model.UserDao
	count := 0
	if err := Dao.DB.Model(&model.UserDao{}).Where("id=?", id).First(&user).Count(&count).Error; err != nil {
		return model.Response{
			Status: e.ErrorDatabase,
			Msg:    "数据库错误",
		}
	}
	if count == 0 {
		return model.Response{
			Status: e.ErrorNotExistUser,
			Msg:    "不存在此用户",
		}
	}
	user.Status = 114514
	Dao.DB.Save(&user)
	return model.Response{
		Status: e.SUCCESS,
		Data:   model.BuildUser(user),
		Msg:    "添加管理员成功",
	}
}
