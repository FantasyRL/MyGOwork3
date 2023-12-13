package adminHandler

import (
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service/AdminService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

// @Summary ListUsers
// @Description ListUsers
// @Accept json/form
// @Produce json
// @Param page_num query int false "页码"
// @Param page_size query int false "每页所展示清单数量"
// @Success 200 {object} model.Response "成功"
// @Failure 400 {object} model.ErrorResponse "请求错误"
// @Failure 500 {object} model.ErrorResponse "内部错误"
// @Router /admin/listusers [POST]
func ListUsers(ctx context.Context, c *app.RequestContext) {
	if !adminLogin(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorAuth,
			Msg:    "你没有访问权限",
		})
		return
	}
	var listUsers AdminService.ListUsersService
	if err := c.BindAndValidate(&listUsers); err != nil {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ERROR,
		})
	} else {
		res := listUsers.List()
		c.JSON(e.SUCCESS, res)
	}
}
