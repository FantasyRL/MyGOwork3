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

// @Summary CheckTask
// @Description CheckTask
// @Accept json/form
// @Produce json
// @Param id path string true "用户id"
// @Param tid path string true "文章id"
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} model.ErrorResponse "请求错误"
// @Router /auth/{id}/task/{tid} [GET]
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
