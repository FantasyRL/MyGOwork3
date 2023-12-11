package adminHandler

import (
	"ToDoList/model"
	"ToDoList/pkg/e"
	"ToDoList/service/AdminService"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func ListUsers(ctx context.Context, c *app.RequestContext) {
	if !adminLogin(ctx, c) {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ErrorAuth,
			Msg:    "你没有访问权限",
		})
		return
	}
	var listUsers AdminService.ListUsers
	if err := c.BindAndValidate(&listUsers); err != nil {
		c.JSON(e.InvalidParams, model.ErrorResponse{
			Status: e.ERROR,
		})
	} else {
		res := listUsers.List()
		c.JSON(e.SUCCESS, res)
	}
}
