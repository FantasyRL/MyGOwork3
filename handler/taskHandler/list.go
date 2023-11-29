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

func ListTask(ctx context.Context, c *app.RequestContext) {
	if !userHandler.IdTask(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorId,
			Msg:    "你没有访问权限",
		})
		return
	}
	var listTask TaskService.ListTaskService
	if err := c.BindAndValidate(&listTask); err != nil {
		c.JSON(e.InvalidParams, err)
	} else {
		id := c.Param("id")
		uid, _ := strconv.Atoi(id)
		res := listTask.List(uint(uid))
		c.JSON(e.SUCCESS, res)
	}
}
