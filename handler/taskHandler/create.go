package taskHandler

import (
	"ToDoList/handler/userHandler"
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service/TaskService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
)

func CreateTask(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	if !userHandler.IdTask(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorId,
			Msg:    "你没有访问权限",
		})
		return
	}
	var createTask TaskService.CreateTaskService
	if err := c.BindAndValidate(&createTask); err != nil {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ERROR,
		})
	} else {
		uid, _ := strconv.Atoi(id)
		res := createTask.Create(uint(uid))
		c.JSON(e.SUCCESS, res)
	}

}
