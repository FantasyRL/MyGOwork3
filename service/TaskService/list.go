package TaskService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func (service *ListTaskService) List(uid uint) model.Response {
	var tasks []model.TaskDao
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if err := Dao.DB.Model(&model.TaskDao{}).Where("uid = ?", uid).Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks).Error; err != nil {
		return model.Response{
			Status: e.ErrorList,
			Msg:    "查询代办清单错误",
		}
	}
	return model.Response{
		Status: e.SUCCESS,
		Data: utils.H{
			"DataList": model.BuildTasks(tasks),
			"total":    count,
		},
		Msg: "清单列表查询成功",
	}
}
