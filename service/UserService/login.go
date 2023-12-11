package UserService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
)

func (service *UserService) Login() model.Response {
	var user model.UserDao
	count := 0
	Dao.DB.Model(&model.UserDao{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 0 {
		return model.Response{
			Status: e.ErrorNotExistUser,
			Msg:    "用户不存在",
		}
	}
	if user.CheckPassword(service.Password) == false {
		return model.Response{
			Status: e.ErrorPassword,
			Msg:    "密码错误",
		}
	}
	if user.Status == 1 {
		return model.Response{
			Status: e.ERROR,
			Msg:    "用户被封禁",
		}
	}
	return model.Response{
		Status: e.SUCCESS,
		Data:   model.BuildUser(user),
		Msg:    "登录成功",
	}
}
