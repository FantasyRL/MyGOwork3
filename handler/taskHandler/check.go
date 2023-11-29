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

func CheckTask(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	txtId := c.Param("tid")
	if !userHandler.IdTask(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorId,
			Msg:    "你没有访问权限",
		})
		return
	}
	var checkTask TaskService.CheckTaskService
	uid, _ := strconv.Atoi(id)
	tid, _ := strconv.Atoi(txtId)
	res := checkTask.Check(uint(uid), tid)
	c.JSON(e.SUCCESS, res)
}
