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

// @Summary CreateTask
// @Description CreateTask
// @Accept json/form
// @Produce json
// @Param id path string true "用户id"
// @Param title query string true "标题"
// @Param content query string true "正文"
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} model.ErrorResponse "请求错误"
// @Router /auth/{id}/task/create [POST]
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
