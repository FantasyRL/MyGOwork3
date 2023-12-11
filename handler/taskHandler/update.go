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

// @Summary UpdateTask
// @Description UpdateTask
// @Accept json/form
// @Produce json
// @Param title string true "标题"
// @Param content string true "正文"
// @Param status int false "状态"
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} e.InvalidParams "请求错误"
// @Failure 500 {object} e.ERROR "内部错误"
// @Router /auth/:id/task/:tid [PUT]
func UpdateTask(ctx context.Context, c *app.RequestContext) {
	id := c.Param("id")
	txtId := c.Param("tid")
	if !userHandler.IdTask(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorId,
			Msg:    "你没有访问权限",
		})
		return
	}
	var updateTask TaskService.UpdateTaskService
	uid, _ := strconv.Atoi(id)
	tid, _ := strconv.Atoi(txtId)
	if err := c.BindAndValidate(&updateTask); err != nil {

		c.JSON(e.InvalidParams, err)
	} else {
		res := updateTask.Update(uint(uid), tid)
		c.JSON(e.SUCCESS, res)
	}

}
