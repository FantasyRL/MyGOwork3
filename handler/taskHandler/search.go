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

// @Summary SearchTask
// @Description SearchTask
// @Accept json/form
// @Produce json
// @Param information true "查询字段"
// @Param page_num int false "页码"
// @Param page_size int false "每页所展示清单数量"
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} e.InvalidParams "请求错误"
// @Failure 500 {object} e.ERROR "内部错误"
// @Router /auth/:id/task/search [POST]
func SearchTask(ctx context.Context, c *app.RequestContext) {
	if !userHandler.IdTask(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorId,
			Msg:    "你没有访问权限",
		})
		return
	}
	var searchTask TaskService.SearchTaskService
	if err := c.BindAndValidate(&searchTask); err != nil {
		c.JSON(e.InvalidParams, err)
	} else {
		id := c.Param("id")
		uid, _ := strconv.Atoi(id)
		res := searchTask.Search(uint(uid))
		c.JSON(e.SUCCESS, res)
	}
}
