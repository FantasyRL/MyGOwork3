package TaskService

import (
	"ToDoList/Dao"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"time"
)

func (service *CreateTaskService) Create(id uint) model.Response {
	var user model.UserDao
	Dao.DB.First(&user, id)
	task := model.TaskDao{
		UserName:  user.UserName,
		Uid:       user.ID,
		Title:     service.Title,
		Status:    0,
		Content:   service.Content,
		StartTime: time.Now().Format(time.RFC3339),
		EndTime:   "null",
		Tid:       user.TidCount + 1,
	}
	if err := Dao.DB.Create(&task).Error; err != nil {
		return model.Response{
			Status: e.ErrorCreate,
			Msg:    "创建备忘录失败",
		}
	}
	user.TidCount++
	Dao.DB.Save(&user)
	return model.Response{
		Status: e.SUCCESS,
		Msg:    "创建成功",
		Data:   model.BuildTask(task),
	}
}
