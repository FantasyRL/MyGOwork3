package UserService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
)

func (service *UserService) Register() model.Response {
	var user model.UserDao
	count := 0
	Dao.DB.Model(&model.UserDao{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	if count == 1 {
		return model.Response{
			Status: e.ErrorExistUser,
			Msg:    "用户已存在",
		}
	}
	user.UserName = service.UserName
	if err := user.SetPassword(service.Password); err != nil {
		return model.Response{
			Status: e.ErrorFailEncryption,
			Msg:    err.Error(),
		}
	}
	if err := Dao.DB.Create(&user).Error; err != nil {
		return model.Response{
			Status: e.ErrorDatabase,
			Msg:    err.Error(),
		}
	}
	return model.Response{
		Status: e.SUCCESS,
		Msg:    "用户创建成功",
	}
}
