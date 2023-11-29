package TaskService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
)

func (service *DeleteTaskService) Delete(uid uint, tid int) model.Response {
	var task model.TaskDao
	count := 0
	Dao.DB.Model(&model.TaskDao{}).Where("uid = ? AND tid = ? ", uid, tid).Count(&count).Find(&task)
	if count == 0 {
		return model.Response{
			Status: e.ErrorCheck,
			Msg:    "不存在此代办",
		}
	}
	if err := Dao.DB.Delete(&task).Error; err != nil {
		return model.Response{
			Status: e.ErrorDelete,
			Msg:    "删除失败",
		}
	}
	return model.Response{
		Status: e.SUCCESS,
		Msg:    "删除成功",
	}
}
