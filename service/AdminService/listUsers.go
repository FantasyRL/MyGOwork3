package AdminService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func (service *ListUsersService) List() model.Response {
	var users []model.UserDao
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if err := Dao.DB.Model(&users).Count(&count).Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&users).Error; err != nil {
		return model.Response{
			Status: e.ErrorDatabase,
			Msg:    "用户列表出错",
		}
	}
	return model.Response{
		Status: e.SUCCESS,
		Data: utils.H{
			"users": users,
			"count": count,
		},
		Msg: "用户列表查询成功",
	}
}
