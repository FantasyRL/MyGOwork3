package TaskService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func (service *SearchTaskService) Search(uid uint) model.Response {
	var tasks []model.TaskDao
	count := 0
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	if err := Dao.DB.Model(&model.TaskDao{}).Where("uid = ?", uid).
		Where("title LIKE ? OR content LIKE ?", "%"+service.Information+"%", "%"+service.Information+"%").Count(&count).
		Limit(service.PageSize).Offset((service.PageNum - 1) * service.PageSize).Find(&tasks).Error; err != nil {
		return model.Response{
			Status: e.ErrorSearch,
			Msg:    "查找失败",
		}
	}
	var Tasks []model.Task
	for _, task := range tasks {
		task.View++
		Dao.DB.Save(&task)
		Tasks = append(Tasks, model.BuildTask(task))
	}
	return model.Response{
		Status: e.SUCCESS,
		Data: utils.H{
			"DataList": Tasks,
			"total":    count,
		},
		Msg: "查找成功",
	}
}
