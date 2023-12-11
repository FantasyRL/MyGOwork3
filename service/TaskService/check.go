package TaskService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
)

func (service *CheckTaskService) Check(uid uint, tid int) model.Response {
	var user model.UserDao
	Dao.DB.First(&user, uid)
	task := model.TaskDao{}
	count := 0
	Dao.DB.Where("uid = ? and tid = ?", uid, tid).Find(&task).Count(&count)
	if count == 0 {
		return model.Response{
			Status: e.ErrorCheck,
			Msg:    "不存在这个tid或已被删除喵",
		}
	} else {
		task.View++
		Dao.DB.Save(&task)
		return model.Response{
			Status: e.SUCCESS,
			Data:   model.BuildTask(task),
			Msg:    "查询成功",
		}
	}
}
