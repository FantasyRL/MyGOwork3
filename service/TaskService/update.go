package TaskService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"time"
)

func (service *UpdateTaskService) Update(uid uint, tid int) model.Response {
	var task model.TaskDao
	if err := Dao.DB.Model(&model.TaskDao{}).Where("uid = ? AND tid = ?", uid, tid).Find(&task).Error; err != nil {
		return model.Response{
			Status: e.ErrorCheck,
			Msg:    "不存在该代办",
		}
	}
	task.Title = service.Title
	task.Content = service.Content
	if task.Status == 0 && service.Status == 1 {
		task.Status = service.Status
		task.EndTime = time.Now().Format(time.DateTime)
	}
	if err := Dao.DB.Save(&task).Error; err != nil {
		return model.Response{
			Status: e.ErrorUpdate,
			Msg:    "更新失败",
		}
	}
	return model.Response{
		Status: e.SUCCESS,
		Data:   model.BuildTask(task),
		Msg:    "更新成功",
	}
}
