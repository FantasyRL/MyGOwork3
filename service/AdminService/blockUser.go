package AdminService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
)

func (service *BlockUserService) Block(id string) model.Response {
	var user model.UserDao
	count := 0
	if err := Dao.DB.Model(&model.UserDao{}).Where("id = ?", id).Count(&count).First(&user).Error; err != nil {
		return model.Response{
			Status: e.ErrorDatabase,
			Msg:    "数据库错误",
		}
	}
	if count == 0 {
		return model.Response{
			Status: e.ErrorNotExistUser,
			Msg:    "用户不存在",
		}
	}
	if user.Status == 0 {
		user.Status = 1
		Dao.DB.Save(&user)
		return model.Response{
			Status: e.SUCCESS,
			Data:   model.BuildUser(user),
			Msg:    "封禁成功",
		}
	}
	return model.Response{
		Status: e.InvalidParams,
		Msg:    "用户已封禁或为管理员",
	}

}
